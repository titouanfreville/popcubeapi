package main

import (
	"github.com/titouanfreville/popcubeapi/api"
	"github.com/titouanfreville/popcubeapi/datastores"
)

func initAPI(ds datastores.DbStore) {
	apiBase := api.Base{}
	apiBase.StartAPI("", "3000", &ds)
}

func initDatastore() datastores.DbStore {
	ds := datastores.DbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev", "0.0.0.0", "3306")
	return ds
}

func main() {
	ds := initDatastore()
	initAPI(ds)
}
