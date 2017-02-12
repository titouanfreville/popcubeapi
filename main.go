package main

import (
	myapi "github.com/titouanfreville/popcubeapi/api"
	datastores "github.com/titouanfreville/popcubeapi/datastores"
)

func api() {
	myapi.API()
}

func init() {
	ds := datastores.DbStore{}
	ds.InitConnection("root", "popcube_test", "popcube_dev")
}

func main() {

}
