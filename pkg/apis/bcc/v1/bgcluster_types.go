package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgClusterSpec defines the desired state of BgCluster
// +k8s:openapi-gen=true
type BgClusterSpec struct {
	Alias        string `json:"alias" protobuf:"bytes,1,opt,name=alias"`
	APIServer    string `json:"apiserver" protobuf:"bytes,2,opt,name=apiserver"`
	MetricServer string `json:"metricserver" protobuf:"bytes,3,opt,name=metricserver"`
	SAToken      string `json:"satoken" protobuf:"bytes,4,opt,name=satoken"`
	// gaoshiyao
	Status       string `json:"status" protobuf:"bytes,5,opt,name=status"`           // 是否在线 up | down
	ClusterType  string `json:"clusterType" protobuf:"bytes,6,opt,name=clusterType"` // 集群类型 core | member
	City         string `json:"city" protobuf:"bytes,7,opt,name=city"`
	ComputerName string `json:"computerName" protobuf:"bytes,8,opt,name=computerName"`
	Node         string `json:"node" protobuf:"bytes,9,opt,name=node"`
	Cpu          string `json:"cpu" protobuf:"bytes,10,opt,name=cpu"`
	Memory       string `json:"memory" protobuf:"bytes,11,opt,name=memory"`
	Fs           string `json:"fs" protobuf:"bytes,12,opt,name=fs"`
	Des          string `json:"des" protobuf:"bytes,13,opt,name=des"`
}

// BgClusterStatus defines the observed state of BgCluster
// +k8s:openapi-gen=true
type BgClusterStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgCluster is the Schema for the bgclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgclusters,scope=Cluster
type BgCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgClusterList contains a list of BgCluster
type BgClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgCluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}
