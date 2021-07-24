package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgRoleUserSpec defines the desired state of BgRoleUser
// +k8s:openapi-gen=true
type BgRoleUserSpec struct {
	Role string `json:"role" protobuf:"bytes,1,opt,name=role"`
	User string `json:"user" protobuf:"bytes,2,opt,name=user"`
}

// BgRoleUserStatus defines the observed state of BgRoleUser
// +k8s:openapi-gen=true
type BgRoleUserStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`

	// 角色编号
	// +optional
	RoleID string `json:"roleid,omitempty" protobuf:"bytes,4,opt,name=roleid"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgRoleUser is the Schema for the bgroleusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgroleusers,scope=Cluster
type BgRoleUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgRoleUserSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgRoleUserStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgRoleUserList contains a list of BgRoleUser
type BgRoleUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgRoleUser `json:"items" protobuf:"bytes,2,rep,name=items"`
}
