package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgGroupSpec defines the desired state of BgGroup
// +k8s:openapi-gen=true
type BgGroupSpec struct {
	Alias   string `json:"alias,omitempty" protobuf:"bytes,1,opt,name=alias"`
	User    string `json:"user,omitempty" protobuf:"bytes,2,opt,name=user"`
	Desc    string `json:"desc,omitempty" protobuf:"bytes,3,opt,name=desc"`
	Default string `json:"default,omitempty" protobuf:"bytes,4,opt,name=default"`
}

// BgGroupStatus defines the observed state of BgGroup
// +k8s:openapi-gen=true
type BgGroupStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,2,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,3,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,4,opt,name=phase,casttype=BgPhase"`

	// 分组下的命名空间
	Namespaces []string `json:"namespaces,omitempty" protobuf:"bytes,5,rep,name=namespaces"`

	// 分组下的用户
	Users []string `json:"users,omitempty" protobuf:"bytes,6,rep,name=users"`

	// The generation observed by the BgGroup controller.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgGroup is the Schema for the bggroups API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bggroups,scope=Cluster
type BgGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgGroupSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgGroupStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgGroupList contains a list of BgGroup
type BgGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgGroup `json:"items" protobuf:"bytes,2,rep,name=items"`
}
