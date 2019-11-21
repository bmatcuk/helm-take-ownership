package pkg

// Many of the ideas here were borrowed from chartify:
// https://github.com/appscode/chartify/blob/5ada66667c74ef0c46c855ade9fce1fe0dfa06ad/pkg/kube_objects.go

import (
	"reflect"
	"strings"

	"github.com/ghodss/yaml"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/api"
	apiref "k8s.io/kubernetes/pkg/api/ref"
)

func (k *Kube) GetKubeObjects(client *kubernetes.Clientset) error {
	k.NumTemplates = 0
	if err := k.getConfigMaps(client); err != nil {
		return err
	}
	if err := k.getDeployments(client); err != nil {
		return err
	}
	if err := k.getDaemonSets(client); err != nil {
		return err
	}
	if err := k.getHorizontalPodAutoscalers(client); err != nil {
		return err
	}
	if err := k.getIngresses(client); err != nil {
		return err
	}
	if err := k.getNetworkPolicies(client); err != nil {
		return err
	}
	if err := k.getJobs(client); err != nil {
		return err
	}
	if err := k.getPersistentVolumeClaims(client); err != nil {
		return err
	}
	if err := k.getPersistentVolumes(client); err != nil {
		return err
	}
	if err := k.getPods(client); err != nil {
		return err
	}
	if err := k.getReplicaSets(client); err != nil {
		return err
	}
	if err := k.getReplicationControllers(client); err != nil {
		return err
	}
	if err := k.getSecrets(client); err != nil {
		return err
	}
	if err := k.getServices(client); err != nil {
		return err
	}
	if err := k.getStatefulSets(client); err != nil {
		return err
	}
	if err := k.getStorageClasses(client); err != nil {
		return err
	}

	return nil
}

func (k *Kube) getConfigMaps(client *kubernetes.Clientset) error {
	for _, name := range k.ConfigMapNames {
		configmap, err := client.CoreV1().ConfigMaps(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, configmap)
		if err != nil {
			return err
		}
		if configmap.Kind == "" {
			configmap.Kind = ref.Kind
		}
		if configmap.APIVersion == "" {
			configmap.APIVersion = ref.APIVersion
		}
		cleanupMeta(&configmap.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(configmap)
		if err != nil {
			return err
		}
		k.ConfigMaps = append(k.ConfigMaps, yml)
	}
	k.NumTemplates += len(k.ConfigMaps)
	return nil
}

func (k *Kube) getDeployments(client *kubernetes.Clientset) error {
	for _, name := range k.DeploymentNames {
		deployment, err := client.ExtensionsV1beta1().Deployments(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, deployment)
		if err != nil {
			return err
		}
		if deployment.Kind == "" {
			deployment.Kind = ref.Kind
		}
		if deployment.APIVersion == "" {
			deployment.APIVersion = makeAPIVersion(deployment.GetSelfLink())
		}
		cleanupMeta(&deployment.ObjectMeta)
		cleanupPodSpec(&deployment.Spec.Template.Spec)
		cleanupDecorators(deployment.ObjectMeta.Annotations)
		deployment.Spec.Selector = nil
		yml, err := cleanupAndMarshalToYaml(deployment)
		if err != nil {
			return err
		}
		k.Deployments = append(k.Deployments, yml)
	}
	k.NumTemplates += len(k.Deployments)
	return nil
}

func (k *Kube) getDaemonSets(client *kubernetes.Clientset) error {
	for _, name := range k.DaemonSetNames {
		daemon, err := client.ExtensionsV1beta1().DaemonSets(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, daemon)
		if err != nil {
			return err
		}
		if daemon.Kind == "" {
			daemon.Kind = ref.Kind
		}
		if daemon.APIVersion == "" {
			daemon.APIVersion = makeAPIVersion(daemon.GetSelfLink())
		}
		cleanupMeta(&daemon.ObjectMeta)
		cleanupPodSpec(&daemon.Spec.Template.Spec)
		yml, err := cleanupAndMarshalToYaml(daemon)
		if err != nil {
			return err
		}
		k.DaemonSets = append(k.DaemonSets, yml)
	}
	k.NumTemplates += len(k.DaemonSets)
	return nil
}

func (k *Kube) getHorizontalPodAutoscalers(client *kubernetes.Clientset) error {
	for _, name := range k.HorizontalPodAutoscalerNames {
		hpa, err := client.AutoscalingV1().HorizontalPodAutoscalers(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, hpa)
		if err != nil {
			return err
		}
		if hpa.Kind == "" {
			hpa.Kind = ref.Kind
		}
		if hpa.APIVersion == "" {
			hpa.APIVersion = makeAPIVersion(hpa.GetSelfLink())
		}
		cleanupMeta(&hpa.ObjectMeta)
		cleanupDecorators(hpa.ObjectMeta.Annotations)
		yml, err := cleanupAndMarshalToYaml(hpa)
		if err != nil {
			return err
		}
		k.HorizontalPodAutoscalers = append(k.HorizontalPodAutoscalers, yml)
	}
	k.NumTemplates += len(k.HorizontalPodAutoscalers)
	return nil
}

func (k *Kube) getIngresses(client *kubernetes.Clientset) error {
	for _, name := range k.IngressNames {
		ingress, err := client.ExtensionsV1beta1().Ingresses(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, ingress)
		if err != nil {
			return err
		}
		if ingress.Kind == "" {
			ingress.Kind = ref.Kind
		}
		if ingress.APIVersion == "" {
			ingress.APIVersion = makeAPIVersion(ingress.GetSelfLink())
		}
		cleanupMeta(&ingress.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(ingress)
		if err != nil {
			return err
		}
		k.Ingresses = append(k.Ingresses, yml)
	}
	k.NumTemplates += len(k.Ingresses)
	return nil
}

func (k *Kube) getNetworkPolicies(client *kubernetes.Clientset) error {
	for _, name := range k.NetworkPolicyNames {
		netpol, err := client.NetworkingV1().NetworkPolicies(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, netpol)
		if err != nil {
			return err
		}
		if netpol.Kind == "" {
			netpol.Kind = ref.Kind
		}
		if netpol.APIVersion == "" {
			netpol.APIVersion = makeAPIVersion(netpol.GetSelfLink())
		}
		cleanupMeta(&netpol.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(netpol)
		if err != nil {
			return err
		}
		k.NetworkPolicies = append(k.NetworkPolicies, yml)
	}
	k.NumTemplates += len(k.NetworkPolicies)
	return nil
}

func (k *Kube) getJobs(client *kubernetes.Clientset) error {
	for _, name := range k.JobNames {
		job, err := client.BatchV1().Jobs(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, job)
		if err != nil {
			return err
		}
		if job.Kind == "" {
			job.Kind = ref.Kind
		}
		if job.APIVersion == "" {
			job.APIVersion = makeAPIVersion(job.GetSelfLink())
		}
		cleanupMeta(&job.ObjectMeta)
		cleanupPodSpec(&job.Spec.Template.Spec)
		cleanupDecorators(job.ObjectMeta.Labels)
		cleanupDecorators(job.Spec.Template.Labels)
		cleanupDecorators(job.Spec.Selector.MatchLabels)
		yml, err := cleanupAndMarshalToYaml(job)
		if err != nil {
			return err
		}
		k.Jobs = append(k.Jobs, yml)
	}
	k.NumTemplates += len(k.Jobs)
	return nil
}

func (k *Kube) getPersistentVolumeClaims(client *kubernetes.Clientset) error {
	for _, name := range k.PersistentVolumeClaimNames {
		pvc, err := client.CoreV1().PersistentVolumeClaims(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, pvc)
		if err != nil {
			return err
		}
		if pvc.Kind == "" {
			pvc.Kind = ref.Kind
		}
		if pvc.APIVersion == "" {
			pvc.APIVersion = ref.APIVersion
		}
		cleanupMeta(&pvc.ObjectMeta)
		cleanupDecorators(pvc.ObjectMeta.Annotations)
		yml, err := cleanupAndMarshalToYaml(pvc)
		if err != nil {
			return err
		}
		k.PersistentVolumeClaims = append(k.PersistentVolumeClaims, yml)
	}
	k.NumTemplates += len(k.PersistentVolumeClaims)
	return nil
}

func (k *Kube) getPersistentVolumes(client *kubernetes.Clientset) error {
	for _, name := range k.PersistentVolumeNames {
		pv, err := client.CoreV1().PersistentVolumes().Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, pv)
		if err != nil {
			return err
		}
		if pv.Kind == "" {
			pv.Kind = ref.Kind
		}
		if pv.APIVersion == "" {
			pv.APIVersion = ref.APIVersion
		}
		cleanupMeta(&pv.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(pv)
		if err != nil {
			return err
		}
		k.PersistentVolumes = append(k.PersistentVolumes, yml)
	}
	k.NumTemplates += len(k.PersistentVolumes)
	return nil
}

func (k *Kube) getPods(client *kubernetes.Clientset) error {
	for _, name := range k.PodNames {
		pod, err := client.CoreV1().Pods(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, pod)
		if err != nil {
			return err
		}
		if pod.Kind == "" {
			pod.Kind = ref.Kind
		}
		if pod.APIVersion == "" {
			pod.APIVersion = ref.APIVersion
		}
		cleanupMeta(&pod.ObjectMeta)
		cleanupPodSpec(&pod.Spec)
		yml, err := cleanupAndMarshalToYaml(pod)
		if err != nil {
			return nil
		}
		k.Pods = append(k.Pods, yml)
	}
	k.NumTemplates += len(k.Pods)
	return nil
}

func (k *Kube) getReplicaSets(client *kubernetes.Clientset) error {
	for _, name := range k.ReplicaSetNames {
		rs, err := client.ExtensionsV1beta1().ReplicaSets(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, rs)
		if err != nil {
			return err
		}
		if rs.Kind == "" {
			rs.Kind = ref.Kind
		}
		if rs.APIVersion == "" {
			rs.APIVersion = makeAPIVersion(rs.GetSelfLink())
		}
		cleanupMeta(&rs.ObjectMeta)
		cleanupPodSpec(&rs.Spec.Template.Spec)
		cleanupDecorators(rs.ObjectMeta.Annotations)
		cleanupDecorators(rs.ObjectMeta.Labels)
		cleanupDecorators(rs.Spec.Selector.MatchLabels)
		cleanupDecorators(rs.Spec.Template.ObjectMeta.Labels)
		yml, err := cleanupAndMarshalToYaml(rs)
		if err != nil {
			return err
		}
		k.ReplicaSets = append(k.ReplicaSets, yml)
	}
	k.NumTemplates += len(k.ReplicaSets)
	return nil
}

func (k *Kube) getReplicationControllers(client *kubernetes.Clientset) error {
	for _, name := range k.ReplicationControllerNames {
		rc, err := client.CoreV1().ReplicationControllers(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, rc)
		if err != nil {
			return err
		}
		if rc.Kind == "" {
			rc.Kind = ref.Kind
		}
		if rc.APIVersion == "" {
			rc.APIVersion = ref.APIVersion
		}
		cleanupMeta(&rc.ObjectMeta)
		cleanupPodSpec(&rc.Spec.Template.Spec)
		yml, err := cleanupAndMarshalToYaml(rc)
		if err != nil {
			return err
		}
		k.ReplicationControllers = append(k.ReplicationControllers, yml)
	}
	k.NumTemplates += len(k.ReplicationControllers)
	return nil
}

func (k *Kube) getSecrets(client *kubernetes.Clientset) error {
	for _, name := range k.SecretNames {
		secret, err := client.CoreV1().Secrets(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, secret)
		if err != nil {
			return err
		}
		if secret.Kind == "" {
			secret.Kind = ref.Kind
		}
		if secret.APIVersion == "" {
			secret.APIVersion = ref.APIVersion
		}
		cleanupMeta(&secret.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(secret)
		if err != nil {
			return err
		}
		k.Secrets = append(k.Secrets, yml)
	}
	k.NumTemplates += len(k.Secrets)
	return nil
}

func (k *Kube) getServices(client *kubernetes.Clientset) error {
	for _, name := range k.ServiceNames {
		service, err := client.CoreV1().Services(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, service)
		if err != nil {
			return err
		}
		if service.Kind == "" {
			service.Kind = ref.Kind
		}
		if service.APIVersion == "" {
			service.APIVersion = ref.APIVersion
		}
		cleanupMeta(&service.ObjectMeta)
		service.Spec.ClusterIP = ""
		for i := range service.Spec.Ports {
			service.Spec.Ports[i].NodePort = 0
		}
		yml, err := cleanupAndMarshalToYaml(service)
		if err != nil {
			return err
		}
		k.Services = append(k.Services, yml)
	}
	k.NumTemplates += len(k.Services)
	return nil
}

func (k *Kube) getStatefulSets(client *kubernetes.Clientset) error {
	for _, name := range k.StatefulSetNames {
		statefulset, err := client.AppsV1beta1().StatefulSets(k.Namespace).Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, statefulset)
		if err != nil {
			return err
		}
		if statefulset.Kind == "" {
			statefulset.Kind = ref.Kind
		}
		if statefulset.APIVersion == "" {
			statefulset.APIVersion = makeAPIVersion(statefulset.GetSelfLink())
		}
		cleanupMeta(&statefulset.ObjectMeta)
		cleanupPodSpec(&statefulset.Spec.Template.Spec)
		yml, err := cleanupAndMarshalToYaml(statefulset)
		if err != nil {
			return err
		}
		k.StatefulSets = append(k.StatefulSets, yml)
	}
	k.NumTemplates += len(k.StatefulSets)
	return nil
}

func (k *Kube) getStorageClasses(client *kubernetes.Clientset) error {
	for _, name := range k.StorageClassNames {
		sc, err := client.StorageV1().StorageClasses().Get(name, metav1.GetOptions{})
		if err != nil {
			return err
		}
		ref, err := apiref.GetReference(api.Scheme, sc)
		if err != nil {
			return err
		}
		if sc.Kind == "" {
			sc.Kind = ref.Kind
		}
		if sc.APIVersion == "" {
			sc.APIVersion = makeAPIVersion(sc.GetSelfLink())
		}
		cleanupMeta(&sc.ObjectMeta)
		yml, err := cleanupAndMarshalToYaml(sc)
		if err != nil {
			return err
		}
		k.StorageClasses = append(k.StorageClasses, yml)
	}
	k.NumTemplates += len(k.StorageClasses)
	return nil
}

func makeAPIVersion(selfLink string) string {
	str := strings.Split(selfLink, "/")
	if len(str) > 2 {
		return (str[2] + "/" + str[3])
	}
	return ""
}

func cleanupMeta(m *metav1.ObjectMeta) {
	m.GenerateName = ""
	m.SelfLink = ""
	m.UID = types.UID("")
	m.ResourceVersion = ""
	m.Generation = 0
	m.CreationTimestamp = metav1.Time{}
	m.DeletionTimestamp = nil
}

func cleanupPodSpec(p *corev1.PodSpec) {
	p.DNSPolicy = corev1.DNSPolicy("")
	p.NodeName = ""
	p.TerminationGracePeriodSeconds = nil
	if p.ServiceAccountName == "default" {
		p.ServiceAccountName = ""
	}
	for i := range p.Containers {
		p.Containers[i].TerminationMessagePath = ""
	}
	for i := range p.InitContainers {
		p.InitContainers[i].TerminationMessagePath = ""
	}
}

func cleanupDecorators(m map[string]string) {
	delete(m, "autoscaling.alpha.kubernetes.io/conditions")
	delete(m, "autoscaling.alpha.kubernetes.io/current-metrics")
	delete(m, "controller-uid")
	delete(m, "deployment.kubernetes.io/desired-replicas")
	delete(m, "deployment.kubernetes.io/max-replicas")
	delete(m, "deployment.kubernetes.io/revision")
	delete(m, "kubectl.kubernetes.io/last-applied-configuration")
	delete(m, "kubernetes.io/change-cause")
	delete(m, "pod-template-hash")
	delete(m, "pv.kubernetes.io/bind-completed")
	delete(m, "pv.kubernetes.io/bound-by-controller")
}

func cleanupAndMarshalToYaml(obj interface{}) (string, error) {
	// marshal to yaml
	yamlData, err := yaml.Marshal(obj)
	if err != nil {
		return "", err
	}

	// unmarshal to a map
	var resource map[string]interface{}
	if err := yaml.Unmarshal(yamlData, &resource); err != nil {
		return "", err
	}
	delete(resource, "status")
	removeEmptyValues(resource)

	// marshal back to yaml
	yamlData, err = yaml.Marshal(resource)
	if err != nil {
		return "", err
	}
	return string(yamlData), nil
}

func removeEmptyValues(resource map[string]interface{}) error {
	for k, v := range resource {
		valueOf := reflect.ValueOf(v)
		if valueOf.Kind() == reflect.Ptr && !valueOf.IsNil() {
			valueOf = valueOf.Elem()
		}
		if !valueOf.IsValid() || isEmpty(valueOf) {
			delete(resource, k)
		} else if valueOf.Kind() == reflect.Map || valueOf.Kind() == reflect.Struct {
			data, err := yaml.Marshal(v)
			if err != nil {
				return err
			}

			var newResource map[string]interface{}
			if err := yaml.Unmarshal(data, &newResource); err != nil {
				return err
			}
			if err := removeEmptyValues(newResource); err != nil {
				return err
			}

			resource[k] = newResource
		} else if valueOf.Kind() == reflect.Array || valueOf.Kind() == reflect.Slice {
			newResource, err := removeEmptyValuesFromSlice(valueOf)
			if err != nil {
				return err
			}
			resource[k] = newResource
		}
	}
	return nil
}

func removeEmptyValuesFromSlice(resource reflect.Value) ([]interface{}, error) {
	newResource := make([]interface{}, resource.Len())
	idx := 0
	for i := 0; i < resource.Len(); i++ {
		valueOf := resource.Index(i)
		if valueOf.Kind() == reflect.Ptr && !valueOf.IsNil() {
			valueOf = valueOf.Elem()
		}
		if !valueOf.IsValid() || isEmpty(valueOf) {
			continue
		}

		if valueOf.Kind() == reflect.Map || valueOf.Kind() == reflect.Struct {
			data, err := yaml.Marshal(valueOf.Interface())
			if err != nil {
				return newResource, err
			}

			var newMapResource map[string]interface{}
			if err := yaml.Unmarshal(data, &newMapResource); err != nil {
				return newResource, err
			}
			if err := removeEmptyValues(newMapResource); err != nil {
				return newResource, err
			}

			newResource[idx] = newMapResource
		} else if valueOf.Kind() == reflect.Array || valueOf.Kind() == reflect.Slice {
			newSliceResource, err := removeEmptyValuesFromSlice(valueOf)
			if err != nil {
				return newResource, err
			}
			newResource[idx] = newSliceResource
		} else {
			newResource[idx] = valueOf.Interface()
		}
		idx++
	}
	return newResource[:idx], nil
}

func isEmpty(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
