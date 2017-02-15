package main

import (
	"github.com/titouanfreville/popcubeapi/api"
	"github.com/titouanfreville/popcubeapi/datastores"
)

func initAPI() {
	apiBase := api.Base{}
	apiBase.StartAPI("", "3000")
}

func initDatastore() {
	ds := datastores.DbStore{}
	ds.InitDatabase("root", "popcube_test", "popcube_dev", "0.0.0.0", "")
	// return ds
}

func main() {
	initDatastore()
	initAPI()
}
