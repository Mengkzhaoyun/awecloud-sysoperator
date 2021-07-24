package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgAppDeploySpec defines the desired state of BgAppDeploy
type BgAppDeploySpec struct {
	User       string `json:"user" protobuf:"bytes,1,opt,name=user"`             // 用户
	App        string `json:"app" protobuf:"bytes,2,opt,name=app"`               // 部署的应用id
	DeployName string `json:"deployName" protobuf:"bytes,3,opt,name=deployName"` // 部署应用名称
	Group      string `json:"group" protobuf:"bytes,4,opt,name=group"`           // 分组
	Notes      string `json:"notes" protobuf:"bytes,5,opt,name=notes"`           // 部署信息
}

// BgAppDeployStatus defines the observed state of BgAppDeploy
type BgAppDeployStatus struct {
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,1,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,2,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgAppDeploy is the Schema for the bgappdeploys API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgappdeploys,scope=Namespaced
type BgAppDeploy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgAppDeploySpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgAppDeployStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgAppDeployList contains a list of BgAppDeploy
type BgAppDeployList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgAppDeploy `json:"items" protobuf:"bytes,2,rep,name=items"`
}
