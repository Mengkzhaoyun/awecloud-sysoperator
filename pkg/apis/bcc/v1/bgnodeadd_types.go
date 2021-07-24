package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgNodeAddSpec defines the desired state of BgNodeAdd
type BgNodeAddSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Image string `json:"image" protobuf:"bytes,1,opt,name=image"`
	// 集群
	CreateUser   string `json:"createUser" protobuf:"bytes,2,opt,name=createUser"`     // 创建用户
	ClusterName  string `json:"clusterName" protobuf:"bytes,3,opt,name=clusterName"`   // 集群名称
	City         string `json:"city" protobuf:"bytes,4,opt,name=city"`                 // 城市
	ComputerRoom string `json:"computerRoom" protobuf:"bytes,5,opt,name=computerRoom"` // 机房
	CreateStatus string `json:"createStatus" protobuf:"bytes,6,opt,name=createStatus"` // 创建状态
	Hosts        string `json:"hosts" protobuf:"bytes,7,opt,name=hosts"`               // hosts 文件
	Systech      string `json:"systech" protobuf:"bytes,8,opt,name=systech"`           // 配置脚本
}

// BgNodeAddStatus defines the observed state of BgNodeAdd
type BgNodeAddStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,1,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,2,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgNodeAdd is the Schema for the bgnodeadds API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgnodeadds,scope=Cluster
type BgNodeAdd struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgNodeAddSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgNodeAddStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgNodeAddList contains a list of BgNodeAdd
type BgNodeAddList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgNodeAdd `json:"items" protobuf:"bytes,2,rep,name=items"`
}
