package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgMenuSpec defines the desired state of BgMenu
type BgMenuSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Alias     string       `json:"alias,omitempty" protobuf:"bytes,1,opt,name=alias"`
	Roles     []string     `json:"roles,omitempty" protobuf:"bytes,2,rep,name=roles"`
	Href      string       `json:"href,omitempty" protobuf:"bytes,3,opt,name=href"`
	Icon      string       `json:"icon,omitempty" protobuf:"bytes,4,opt,name=icon"`
	Iconhov   string       `json:"iconhov,omitempty" protobuf:"bytes,5,opt,name=iconhov"`
	Group     bool         `json:"group,omitempty" protobuf:"varint,6,opt,name=group"`
	Namespace bool         `json:"namespace,omitempty" protobuf:"varint,7,opt,name=namespace"`
	IsNewTab  bool         `json:"is_new_tab,omitempty" protobuf:"varint,8,opt,name=is_new_tab,json=isNewTab"`
	Submenus  []BgMenuSpec `json:"submenus,omitempty" protobuf:"bytes,9,rep,name=submenus"`
}

// type Menu struct {
// 	Alias     string   `json:"alias,omitempty"`
// 	Roles     []string `json:"roles,omitempty"`
// 	Href      string   `json:"href,omitempty"`
// 	Icon      string   `json:"icon,omitempty"`
// 	Iconhov   string   `json:"iconhov,omitempty"`
// 	Group     bool     `json:"group,omitempty"`
// 	Namespace bool     `json:"namespace,omitempty"`
// 	IsNewTab  bool     `json:"is_new_tab,omitempty"`
// 	Submenus  []Menu   `json:"submenus,omitempty"`
// }

//type Submenu struct {
//	Alias     string   `json:"alias"`
//	Href      string   `json:"href"`
//	Icon      string   `json:"icon"`
//	Iconhov   string   `json:"iconhov"`
//	Roles     []string `json:"roles"`
//	Group     bool     `json:"group"`
//	Namespace bool     `json:"namespace"`
//}

// BgMenuStatus defines the observed state of BgMenu
type BgMenuStatus struct {
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

// BgMenu is the Schema for the bgmenus API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgmenus,scope=Cluster
type BgMenu struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgMenuSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgMenuStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgMenuList contains a list of BgMenu
type BgMenuList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgMenu `json:"items" protobuf:"bytes,2,rep,name=items"`
}
