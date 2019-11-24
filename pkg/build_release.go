package pkg

// Helm Release objects: https://github.com/kubernetes/helm/tree/v2.7.2/pkg/proto/hapi

import (
	"fmt"

	chart "k8s.io/helm/pkg/proto/hapi/chart"
	hapi "k8s.io/helm/pkg/proto/hapi/release"
	"k8s.io/helm/pkg/timeconv"
)

func (k *Kube) BuildRelease(name string) *hapi.Release {
	templates, manifest := k.buildTemplates()
	return &hapi.Release{
		Name:      name,
		Info:      k.buildReleaseInfo(),
		Chart:     k.buildReleaseChart(templates),
		Manifest:  manifest,
		Version:   1,
		Namespace: k.Namespace,
	}
}

func (k *Kube) buildTemplates() ([]*chart.Template, string) {
	templates := make([]*chart.Template, k.NumTemplates)
	manifest := ""
	idx := 0
	for i, name := range k.ConfigMapNames {
		manifest += addChartTemplate(templates, "configmap", name, k.ConfigMaps[i], idx)
		idx++
	}
	for i, name := range k.DaemonSetNames {
		manifest += addChartTemplate(templates, "daemonset", name, k.DaemonSets[i], idx)
		idx++
	}
	for i, name := range k.DeploymentNames {
		manifest += addChartTemplate(templates, "deployment", name, k.Deployments[i], idx)
		idx++
	}
	for i, name := range k.HorizontalPodAutoscalerNames {
		manifest += addChartTemplate(templates, "hpa", name, k.HorizontalPodAutoscalers[i], idx)
		idx++
	}
	for i, name := range k.IngressNames {
		manifest += addChartTemplate(templates, "ingress", name, k.Ingresses[i], idx)
		idx++
	}
	for i, name := range k.NetworkPolicyNames {
		manifest += addChartTemplate(templates, "networkpolicy", name, k.NetworkPolicies[i], idx)
		idx++
	}
	for i, name := range k.JobNames {
		manifest += addChartTemplate(templates, "job", name, k.Jobs[i], idx)
		idx++
	}
	for i, name := range k.PersistentVolumeClaimNames {
		manifest += addChartTemplate(templates, "pvc", name, k.PersistentVolumeClaims[i], idx)
		idx++
	}
	for i, name := range k.PersistentVolumeNames {
		manifest += addChartTemplate(templates, "pv", name, k.PersistentVolumes[i], idx)
		idx++
	}
	for i, name := range k.PodNames {
		manifest += addChartTemplate(templates, "pod", name, k.Pods[i], idx)
		idx++
	}
	for i, name := range k.ReplicaSetNames {
		manifest += addChartTemplate(templates, "replicaset", name, k.ReplicaSets[i], idx)
		idx++
	}
	for i, name := range k.ReplicationControllerNames {
		manifest += addChartTemplate(templates, "replicationcontroller", name, k.ReplicationControllers[i], idx)
		idx++
	}
	for i, name := range k.SecretNames {
		manifest += addChartTemplate(templates, "secret", name, k.Secrets[i], idx)
		idx++
	}
	for i, name := range k.ServiceNames {
		manifest += addChartTemplate(templates, "service", name, k.Services[i], idx)
		idx++
	}
	for i, name := range k.StatefulSetNames {
		manifest += addChartTemplate(templates, "statefulset", name, k.StatefulSets[i], idx)
		idx++
	}
	for i, name := range k.StorageClassNames {
		manifest += addChartTemplate(templates, "storageclass", name, k.StorageClasses[i], idx)
		idx++
	}
	return templates, manifest
}

func addChartTemplate(templates []*chart.Template, templateType, name, template string, idx int) string {
	filename := fmt.Sprintf("templates/%v-%v.yaml", templateType, name)
	templates[idx] = &chart.Template{
		Name: filename,
		Data: []byte(template),
	}
	return fmt.Sprintf("\n---\n# Source: %v\n%v", filename, template)
}

func (k *Kube) buildReleaseInfo() *hapi.Info {
	status := &hapi.Status{
		Code: hapi.Status_DEPLOYED,
	}

	info := &hapi.Info{
		Status:        status,
		FirstDeployed: timeconv.Now(),
		LastDeployed:  timeconv.Now(),
		Description:   "Transferred ownership to Helm via helm-take-ownership",
	}

	return info
}

func (k *Kube) buildReleaseChart(templates []*chart.Template) *chart.Chart {
	return &chart.Chart{
		Metadata: &chart.Metadata{
			Name:        k.ChartName,
			Version:     k.ChartVersion,
			Description: "Chart built by helm-take-ownership",
			ApiVersion:  "v1",
		},
		Templates: templates,
	}
}
