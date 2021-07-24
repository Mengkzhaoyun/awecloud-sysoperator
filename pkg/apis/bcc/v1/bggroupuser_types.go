package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgGroupUserSpec defines the desired state of BgGroupUser
// +k8s:openapi-gen=true
type BgGroupUserSpec struct {
	Group string `json:"group" protobuf:"bytes,1,opt,name=group"`
	User  string `json:"user" protobuf:"bytes,2,opt,name=user"`
	Role  string `json:"role" protobuf:"bytes,3,opt,name=role"`
	// 操作状态， Running , Completed
	State BgPhase `json:"state,omitempty" protobuf:"bytes,4,opt,name=state,casttype=BgPhase"`
}

// BgGroupUserStatus defines the observed state of BgGroupUser
// +k8s:openapi-gen=true
type BgGroupUserStatus struct {
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

// BgGroupUser is the Schema for the bggroupusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bggroupusers,scope=Cluster
type BgGroupUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgGroupUserSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgGroupUserStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgGroupUserList contains a list of BgGroupUser
type BgGroupUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgGroupUser `json:"items" protobuf:"bytes,2,rep,name=items"`
}
