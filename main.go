package main

import (
	"github.com/titouanfreville/popcubeapi/api"
	"github.com/titouanfreville/popcubeapi/datastores"
)

func initAPI() {
	api.StartAPI("", "3000")
}

func initDatastore() {
	datastores.Store().InitDatabase("root", "popcube_test", "popcube_dev", "database", "3306")
}

func main() {
	initDatastore()
	initAPI()
}
