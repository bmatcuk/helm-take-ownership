package pkg

type Kube struct {
	ConfigMapNames             []string
	DeploymentNames            []string
	DaemonSetNames             []string
	JobNames                   []string
	PersistentVolumeNames      []string
	PersistentVolumeClaimNames []string
	PodNames                   []string
	ReplicaSetNames            []string
	ReplicationControllerNames []string
	SecretNames                []string
	ServiceNames               []string
	StatefulSetNames           []string
	StorageClassNames          []string

	NumTemplates           int
	ConfigMaps             []string
	Deployments            []string
	DaemonSets             []string
	Jobs                   []string
	PersistentVolumes      []string
	PersistentVolumeClaims []string
	Pods                   []string
	ReplicaSets            []string
	ReplicationControllers []string
	Secrets                []string
	Services               []string
	StatefulSets           []string
	StorageClasses         []string

	ChartName    string
	ChartVersion string
	Namespace    string
}
