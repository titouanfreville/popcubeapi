package main

import (
	api "github.com/titouanfreville/popcubeapi/api"
	datastores "github.com/titouanfreville/popcubeapi/datastores"
)

func initAPI() {
	api.StartAPI("", "3000")
}

func initDatastore() {
	ds := datastores.DbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
}

func main() {
	initAPI()
	initDatastore()
}
