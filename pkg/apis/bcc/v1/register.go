// NOTE: Boilerplate only.  Ignore this file.

// Package v1 contains API Schema definitions for the bcc v1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=bcc.bd-apaas.com
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "bcc.bd-apaas.com", Version: "v1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)
)

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BgUser{},
		&BgUserList{},
		&BgToken{},
		&BgTokenList{},
		&BgStorage{},
		&BgStorageList{},
		&BgSetting{},
		&BgSettingList{},
		&BgRoleUser{},
		&BgRoleUserList{},
		&BgNodeAdd{},
		&BgNodeAddList{},
		&BgNamespace{},
		&BgNamespaceList{},
		&BgMenu{},
		&BgMenuList{},
		&BgLoginLog{},
		&BgLoginLogList{},
		&BgGroupUser{},
		&BgGroupUserList{},
		&BgGroupNamespace{},
		&BgGroupNamespaceList{},
		&BgGroup{},
		&BgGroupList{},
		&BgCluster{},
		&BgClusterList{},
		&BgCloudShell{},
		&BgCloudShellList{},
		&BgAppStore{},
		&BgAppStoreList{},
		&BgAppDeploy{},
		&BgAppDeployList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
