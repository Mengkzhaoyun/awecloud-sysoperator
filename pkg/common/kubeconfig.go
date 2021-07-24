/*
 * @Author: gaoshiyao
 * @Date: 2021-07-22 22:12:22
 * @Description:
 * @LastEditors: gaoshiyao
 * @LastEditTime: 2021-07-22 22:24:00
 */
package common

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

var SaToken string

type Config struct {
	Name   string `json:"name"`
	Server string `json:"server"`
	Token  string `json:"token"`
}

type KubeConfig struct {
	ApiVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	Clusters   []Cluster `json:"clusters"`
	Contexts   []Context `json:"contexts"`
	Users      []User    `json:"users"`
}

type User struct {
	Name string `json:"name"`
	User struct {
		Token string `json:"token"`
	}
}

type Context struct {
	Name    string `json:"name"`
	Context struct {
		Cluster string `json:"cluster"`
		User    string `json:"user"`
	} `json:"context"`
}

type Cluster struct {
	Name    string `json:"name"`
	Cluster struct {
		InsecureSkipTlsVerify bool   `json:"insecure-skip-tls-verify" yaml:"insecure-skip-tls-verify"`
		Server                string `json:"cluster"`
	} `json:"cluster"`
}

//func (c Config) CreateKubeConfigFile() {
//	tmpl, err := template.New("test").Parse(Tmp)
//	if err != nil {
//		panic(err)
//	}
//	//tmpl.Execute(os.Stdout, cup)
//	path := "G:\\GoPath\\src\\gitlab.wodcloud.com\\cloud\\awecloud-sysoperator\\.kube\\" + c.Name
//	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
//	if err != nil {
//		fmt.Println("open failed err:", err)
//		return
//	}
//	tmpl.Execute(file, c)
//}

func (k KubeConfig) ToYamlFile(name string) (err error) {
	out, err := yaml.Marshal(k)
	if err != nil {
		return err
	} else {
		err = ioutil.WriteFile(name, out, 0755)
	}
	return
}

func GetKubePath() (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	kubePath := path.Join(u.HomeDir, ".kube")
	_, err = os.Stat(kubePath)
	if err != nil {
		err = os.Mkdir(kubePath, os.ModePerm)
	}
	return kubePath, err
}

func InitDefaultKubeConfig() (err error) {
	kubePath, err := GetKubePath()
	if err != nil {
		return
	}
	tokenPath := "/var/run/secrets/kubernetes.io/serviceaccount/token"
	b, err := ioutil.ReadFile(tokenPath)
	if err != nil {
		return
	}
	SaToken = string(b)
	config := &Config{"default", "default", SaToken}
	return config.GetDefaultKubeConfig().ToYamlFile(path.Join(kubePath, "default"))
}

func (c Config) GetDefaultKubeConfig() KubeConfig {
	kubeConfig := KubeConfig{ApiVersion: "v1", Kind: "Config"}
	cluster := Cluster{Name: "default"}
	cluster.Cluster.Server = "https://kubernetes.default:443"
	cluster.Cluster.InsecureSkipTlsVerify = true
	kubeConfig.Clusters = append(kubeConfig.Clusters, cluster)
	context := Context{Name: "default"}
	context.Context.Cluster = "default"
	context.Context.User = "default"
	kubeConfig.Contexts = append(kubeConfig.Contexts, context)
	user := User{Name: "default"}
	user.User.Token = c.Token
	kubeConfig.Users = append(kubeConfig.Users, user)
	return kubeConfig
}

func (c Config) CreateKubeConfigFile() (err error) {
	defaultConfig := c.GetDefaultKubeConfig()
	cluster := Cluster{Name: c.Name}
	cluster.Cluster.Server = c.Server
	cluster.Cluster.InsecureSkipTlsVerify = true
	defaultConfig.Clusters = append(defaultConfig.Clusters, cluster)
	context := Context{Name: c.Name}
	context.Context.Cluster = c.Name
	context.Context.User = c.Name
	defaultConfig.Contexts = append(defaultConfig.Contexts, context)
	user := User{Name: c.Name}
	user.User.Token = c.Token
	defaultConfig.Users = append(defaultConfig.Users, user)
	// TODO 创建文件
	kubePath, err := GetKubePath()
	if err != nil {
		return
	}
	return defaultConfig.ToYamlFile(path.Join(kubePath, c.Name))
}
