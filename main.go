package api

import (
	myapi "github.com/titouanfreville/popcubeapi/api"
	ds "github.com/titouanfreville/popcubeapi/datastores"
)

func api() {
	myapi.API()
}

func init() {
	ds.InitConnection()
}
