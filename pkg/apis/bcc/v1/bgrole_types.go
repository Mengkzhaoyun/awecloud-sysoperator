package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// BgRoleSA , 角色预定义-超管
	BgRoleSA = "beagle:kubernetes:sa"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgRoleSpec defines the desired state of BgRole
// +k8s:openapi-gen=true
type BgRoleSpec struct {
	ID    string `json:"id" protobuf:"bytes,1,opt,name=id"`
	Alias string `json:"alias,omitempty" protobuf:"bytes,2,opt,name=alias"`
	Desc  string `json:"desc,omitempty" protobuf:"bytes,3,opt,name=desc"`
}

// BgRoleStatus defines the observed state of BgRole
// +k8s:openapi-gen=true
type BgRoleStatus struct {
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

// BgRole is the Schema for the bgroles API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgroles,scope=Cluster
type BgRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgRoleSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgRoleStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgRoleList contains a list of BgRole
type BgRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgRole `json:"items" protobuf:"bytes,2,rep,name=items"`
}
