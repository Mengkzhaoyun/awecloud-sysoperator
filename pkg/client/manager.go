package client

import (
	"os"

	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
	"sigs.k8s.io/controller-runtime/pkg/client/config"

	"github.com/labstack/gommon/log"
)

// NewClient , new client
func NewClient() client.Client {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Error(err, "pkg.client.NewClientManager() GetConfig Error")
		os.Exit(1)
	}
	mapper, err := apiutil.NewDynamicRESTMapper(cfg)
	if err != nil {
		log.Error(err, "pkg.client.NewClientManager() NewDynamicRESTMapper Error")
		os.Exit(1)
	}
	client, err := client.New(cfg, client.Options{Scheme: scheme.Scheme, Mapper: mapper})
	if err != nil {
		log.Error(err, "pkg.client.NewClientManager() client.New Error")
		os.Exit(1)
	}
	return client
}
