package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/bmatcuk/helm-take-ownership/pkg"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/helm/pkg/storage"
	"k8s.io/helm/pkg/storage/driver"
	"k8s.io/helm/pkg/tiller/environment"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

var (
	version = "master"
	date    = "1970-01-01_00:00:00"
)

func main() {
	var releaseName string
	var printVersion, debug, dryRun bool
	var kubeConfig clientcmd.ClientConfig
	kube := pkg.Kube{}

	helmPluginName := os.Getenv("HELM_PLUGIN_NAME")
	if helmPluginName == "" {
		helmPluginName = "own"
	}

	rootCmd := &cobra.Command{
		Use:     fmt.Sprintf("helm %v RELEASE_NAME", helmPluginName),
		Short:   "Transfer ownership of a k8s release to helm",
		Long:    `The helm-take-ownership plugin transfers ownership from a non-helm-deployed release to helm's management.`,
		Example: fmt.Sprintf("helm %v --deploy my-deployment --svc my-service -n stg my-release", helmPluginName),
		Args: func(c *cobra.Command, args []string) error {
			if !printVersion {
				if len(args) < 1 {
					return errors.New("You must specify the release name.")
				}
				if len(args) > 1 {
					return errors.New("Too many arguments.")
				}
			}
			return nil
		},
		RunE: func(c *cobra.Command, args []string) error {
			if printVersion {
				fmt.Printf("helm-take-ownership v%v (%.10s)", version, date)
				return nil
			}

			releaseName = args[0]
			if kube.ChartName == "RELEASE_NAME" {
				kube.ChartName = releaseName
			}

			var valid bool
			var err error
			kube.Namespace, valid, err = kubeConfig.Namespace()
			if !valid {
				if err != nil {
					return err
				} else {
					return errors.New("You must specify a namespace")
				}
			}

			// create kubernetes factory
			factory := cmdutil.NewFactory(kubeConfig)
			kubernetesClientSet, err := factory.KubernetesClientSet()
			if err != nil {
				return err
			}
			clientset, err := factory.ClientSet()
			if err != nil {
				return err
			}

			// download the kube objects
			log.Print("Downloading kubernetes objects...")
			if err := kube.GetKubeObjects(kubernetesClientSet); err != nil {
				return err
			}

			// create a release
			log.Print("Constructing Helm Release...")
			release := kube.BuildRelease(releaseName)

			// print chart
			if debug {
				data, err := yaml.Marshal(release)
				if err != nil {
					return err
				}
				log.Print("Helm Release:")
				log.Print(string(data))
			}

			// install the release
			if dryRun {
				log.Print("Skipping install because this is a dry run.")
				if debug {
					kubeClientConfig, err := kubeConfig.RawConfig()
					if err != nil {
						return err
					}

					data, err := yaml.Marshal(kubeClientConfig)
					if err != nil {
						return err
					}
					log.Print("Kube config:")
					log.Print(string(data))
				}
			} else {
				log.Print("Installing Helm Chart...")
				cfgmaps := driver.NewConfigMaps(clientset.Core().ConfigMaps(environment.DefaultTillerNamespace))
				cfgmaps.Log = helmPrinter
				releases := storage.Init(cfgmaps)
				releases.Log = helmPrinter
				if err := releases.Create(release); err != nil {
					return err
				}
			}

			return nil
		},
	}

	flags := rootCmd.Flags()
	flags.StringVar(&kube.ChartName, "chartname", "RELEASE_NAME", "Name of the helm Chart")
	flags.StringVar(&kube.ChartVersion, "chartversion", "0.1.0", "Version of the helm Chart")
	flags.BoolVarP(&printVersion, "version", "v", false, "Print version and exit")
	flags.BoolVarP(&debug, "debug", "d", debug, "Debug mode")
	flags.BoolVar(&dryRun, "dry-run", dryRun, "Dry-run")

	flags.StringSliceVar(&kube.ConfigMapNames, "configmaps", kube.ConfigMapNames, "(aka 'cm') Comma-separated names of ConfigMaps to include")
	flags.StringSliceVar(&kube.DaemonSetNames, "daemonsets", kube.DaemonSetNames, "(aka 'ds') Comma-separated names of DaemonSets to include")
	flags.StringSliceVar(&kube.DeploymentNames, "deployments", kube.DeploymentNames, "(aka 'deploy') Comma-separated names of Deployments to include")
	flags.StringSliceVar(&kube.HorizontalPodAutoscalerNames, "horizontalpodautoscalers", kube.HorizontalPodAutoscalers, "(aka 'hpa') Comma-separated names of Horizontal Pod Autoscalers to include")
	flags.StringSliceVar(&kube.IngressNames, "ingresses", kube.IngressNames, "(aka 'ing') Comma-separated names of Ingresses to include")
	flags.StringSliceVar(&kube.NetworkPolicyNames, "networkpolicies", kube.NetworkPolicyNames, "(aka 'netpol') Comma-separated names of NetworkPolicies to include")
	flags.StringSliceVar(&kube.JobNames, "jobs", kube.JobNames, "Comma-separated names of Jobs to include")
	flags.StringSliceVar(&kube.PersistentVolumeClaimNames, "persistentvolumeclaims", kube.PersistentVolumeClaimNames, "(aka 'pvc') Comma-separated names of PersistentVolumeClaims to include")
	flags.StringSliceVar(&kube.PersistentVolumeNames, "persistentvolumes", kube.PersistentVolumeNames, "(aka 'pv') Comma-separated names of PersistentVolumes to include")
	flags.StringSliceVar(&kube.PodNames, "pods", kube.PodNames, "(aka po') Comma-separated names of Pods to include")
	flags.StringSliceVar(&kube.ReplicaSetNames, "replicasets", kube.ReplicaSetNames, "(aka 'rs') Comma-separated names of ReplicaSets to include")
	flags.StringSliceVar(&kube.ReplicationControllerNames, "replicationcontrollers", kube.ReplicationControllerNames, "(aka 'rc') Comma-separated names of ReplicationControllers to include")
	flags.StringSliceVar(&kube.SecretNames, "secrets", kube.SecretNames, "Comma-separated names of Secrets to include")
	flags.StringSliceVar(&kube.ServiceNames, "services", kube.ServiceNames, "(aka 'svc') Comma-separated names of Services to include")
	flags.StringSliceVar(&kube.StatefulSetNames, "statefulsets", kube.StatefulSetNames, "Comma-separated names of StatefulSets to include")
	flags.StringSliceVar(&kube.StorageClassNames, "storageclasses", kube.StorageClassNames, "Comma-separated names of StorageClasses to include")

	// Add generic kubectl flags. Taken from:
	// https://github.com/kubernetes/kubernetes/blob/v1.8.0/pkg/kubectl/cmd/util/factory_client_access.go#L156
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig

	flags.StringVar(&loadingRules.ExplicitPath, "kubeconfig", "", "Path to the kubeconfig file to use for CLI requests.")

	kubeOverrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults}

	flagNames := clientcmd.RecommendedConfigOverrideFlags("k8s-")
	flagNames.ClusterOverrideFlags.APIServer.ShortName = "s"

	clientcmd.BindOverrideFlags(kubeOverrides, flags, flagNames)
	kubeConfig = clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, kubeOverrides)

	// support for 2-3 letter shorthands similar to kubectl
	flags.SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		switch name {
		case "cm":
			name = "configmaps"
		case "ds":
			name = "daemonsets"
		case "deploy":
			name = "deployments"
		case "hpa":
			name = "horizontalpodautoscalers"
		case "ing":
			name = "ingresses"
		case "pvc":
			name = "persistentvolumeclaims"
		case "pv":
			name = "persistentvolumes"
		case "po":
			name = "pods"
		case "rs":
			name = "replicasets"
		case "rc":
			name = "replicationcontrollers"
		case "svc":
			name = "services"
		}
		return pflag.NormalizedName(name)
	})

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func helmPrinter(s string, params ...interface{}) {
	params = append([]interface{}{s}, params...)
	log.Print(params...)
}
