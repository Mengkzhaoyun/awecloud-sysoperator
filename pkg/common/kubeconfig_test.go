package common

import "testing"

func TestName(t *testing.T) {
	//c := Config{}
	//c.Name = "default"
	//c.Server ="https://k8s.wodcloud.com:6443"
	//c.Token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImJDNS1iNVNCanZqZXplZTBCdVVSb0tHNzRjYzBGTkhfWVM1MWRzcnBXeFUifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhd2VjbG91ZC1zeXNvcGVyYXRvci10b2tlbi16OTcycCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhd2VjbG91ZC1zeXNvcGVyYXRvciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6ImJhM2RlMDg4LWJhYWMtNDc5Zi05ZjVjLThmNTMzOTExMzFlOCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTphd2VjbG91ZC1zeXNvcGVyYXRvciJ9.RnRbzRlqkHZGMZ5w5DuQ84rYO12cmfai3TCfYFfSnMuMbcjvPvuPM1jLcmakB3FHbt3C_eC69y217W0MamFL9yz8RVwlLpk5UUxFQGVCcBs6_dQkMBUKr88VH5VCLxAkJCLWMo0iV8qelr5bHEsrDWMhHh0GMURilDlRLq2dMgM9A_T0_XaXQAASgi_O1wm4gy_KPXb1RMWfthRkxlVp7JXCwTYdNEvUeXeO0VoDgDYgrBPZvsSzFw3JgQ74gjn4COuVqzhCQBFxzj15OXR0OMa_3i7cOaU0qqhBXD6ePIF0aTq1rKkWv7aLuBbAR3EYbxeiFcdIxurEpxp045406g"
	//c.CreateKubeConfigFile()

	//c := Config{}
	//c.Name = "ysgz-test"
	//c.Server = "https://103.81.5.56:6443"
	//c.Token = "7176d48e4e66ddb3557a82f2dd316a93"
	////c.CreateKubeConfigFile()
	//c.CreateKubeConfigFile().ToYamlFile()
	if err := InitDefaultKubeConfig(); err != nil {
		panic(err)
	}
}
