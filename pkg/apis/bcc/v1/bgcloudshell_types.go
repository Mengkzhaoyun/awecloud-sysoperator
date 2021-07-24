package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgCloudShellSpec defines the desired state of BgCloudShell
// +k8s:openapi-gen=true
type BgCloudShellSpec struct {
	User      string `json:"user" protobuf:"bytes,1,opt,name=user"`
	Namespace string `json:"namespace" protobuf:"bytes,2,opt,name=namespace"`
}

// BgCloudShellStatus defines the observed state of BgCloudShell
// +k8s:openapi-gen=true
type BgCloudShellStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`

	// RepositoryShell , Pod Namespace
	// +optional
	RepositoryShell string `json:"repositoryshell,omitempty" protobuf:"bytes,4,opt,name=repositoryshell"`

	// RepositoryDocker , Pod Namespace
	// +optional
	RepositoryDocker string `json:"repositorydocker,omitempty" protobuf:"bytes,5,opt,name=repositorydocker"`

	// PodNamespace , Pod Namespace
	// +optional
	PodNamespace string `json:"podnamespace,omitempty" protobuf:"bytes,6,opt,name=podnamespace"`

	// PodName , Pod Name
	// +optional
	PodName string `json:"podname,omitempty" protobuf:"bytes,7,opt,name=podname"`

	// PodContainer , Pod Namespace
	// +optional
	PodContainer string `json:"podcontainer,omitempty" protobuf:"bytes,8,opt,name=podcontainer"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgCloudShell is the Schema for the bgcloudshells API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgcloudshells,scope=Cluster
type BgCloudShell struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgCloudShellSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgCloudShellStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgCloudShellList contains a list of BgCloudShell
type BgCloudShellList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgCloudShell `json:"items" protobuf:"bytes,2,rep,name=items"`
}
