package pkg

type Kube struct {
	ConfigMapNames               []string
	DaemonSetNames               []string
	DeploymentNames              []string
	HorizontalPodAutoscalerNames []string
	IngressNames                 []string
	NetworkPolicyNames           []string
	JobNames                     []string
	PersistentVolumeClaimNames   []string
	PersistentVolumeNames        []string
	PodNames                     []string
	ReplicaSetNames              []string
	ReplicationControllerNames   []string
	SecretNames                  []string
	ServiceNames                 []string
	StatefulSetNames             []string
	StorageClassNames            []string

	NumTemplates             int
	ConfigMaps               []string
	DaemonSets               []string
	Deployments              []string
	HorizontalPodAutoscalers []string
	Ingresses                []string
	NetworkPolicies          []string
	Jobs                     []string
	PersistentVolumeClaims   []string
	PersistentVolumes        []string
	Pods                     []string
	ReplicaSets              []string
	ReplicationControllers   []string
	Secrets                  []string
	Services                 []string
	StatefulSets             []string
	StorageClasses           []string

	ChartName    string
	ChartVersion string
	Namespace    string
}
