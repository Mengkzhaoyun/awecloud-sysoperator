package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BgUser ， 用户对象，仅解决登录的问题
// 说明： 与用户BgUser有关系的对象 ， BgRoleUser ， BgGroupUser
//        移除BgUser时，要移除与之相关的绑定，角色与Group

// BgUserSpec defines the desired state of BgUser
// +k8s:openapi-gen=true
type BgUserSpec struct {
	Alias       string         `json:"alias" protobuf:"bytes,1,opt,name=alias"`
	Phone       string         `json:"phone,omitempty" protobuf:"bytes,2,opt,name=phone"`
	Email       string         `json:"email" protobuf:"bytes,3,opt,name=email"`
	Avatar      string         `json:"avatar,omitempty" protobuf:"bytes,4,opt,name=avatar"`
	Disable     bool           `json:"disable,omitempty" protobuf:"varint,5,opt,name=disable"`
	Password    string         `json:"password,omitempty" protobuf:"bytes,6,opt,name=password"`
	Default     *BgUserDefault `json:"default,omitempty" protobuf:"bytes,7,opt,name=default"`
	CenterAdmin string         `json:"centerAdmin,omitempty" protobuf:"bytes,8,opt,name=centerAdmin"` // 中心管理员开关 off on
	SaAdmin     string         `json:"saAdmin,omitempty" protobuf:"bytes,9,opt,name=saAdmin"`         // 集群管理员开关
}

// BgUserStatus defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`

	// Token ， 会Watch ServiceAccount.Secrets，随之变化
	// +optional
	Token string `json:"token,omitempty" protobuf:"bytes,4,opt,name=token"`

	// 分组
	// +optional
	Groups map[string]BgUserGroup `json:"groups,omitempty" protobuf:"bytes,5,rep,name=groups"`

	// 角色
	// +optional
	Roles []string `json:"roles,omitempty" protobuf:"bytes,6,rep,name=roles"`

	// 用户设置密码以后，密码Password将会从Spec从删除加密保存在Status之中
	// +optional
	Password string `json:"password,omitempty" protobuf:"bytes,7,opt,name=password"`

	// 在用户进入垃圾回收时，设置Disable=true
	// +optional
	Disable bool `json:"disable,omitempty" protobuf:"varint,8,opt,name=disable"`

	// 用户最后一次登录时间
	// +optional
	LastLogin *metav1.Time `json:"lastLogin,omitempty" protobuf:"bytes,9,opt,name=lastLogin"`
}

// BgUserDefault defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserDefault struct {
	Group   string `json:"group,omitempty" protobuf:"bytes,1,opt,name=group"`
	Cluster string `json:"cluster,omitempty" protobuf:"bytes,2,opt,name=cluster"`
}

// BgUserGroup defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserGroup struct {
	Name       string   `json:"name" protobuf:"bytes,1,opt,name=name"`
	Role       string   `json:"role" protobuf:"bytes,2,opt,name=role"`
	Default    string   `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	Namespaces []string `json:"namespaces,omitempty" protobuf:"bytes,4,rep,name=namespaces"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgUser is the Schema for the bgusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgusers,scope=Cluster
type BgUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgUserSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgUserStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgUserList contains a list of BgUser
type BgUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgUser `json:"items" protobuf:"bytes,2,rep,name=items"`
}
