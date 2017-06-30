package main

// import (
// 	"github.com/titouanfreville/popcubeapi/api"
// 	"github.com/titouanfreville/popcubeapi/configs"
// 	"github.com/titouanfreville/popcubeapi/datastores"
// )

// var (
// 	// DbConnectionInfo information to conect to DB
// 	DbConnectionInfo = &configs.DbConnection{}
// 	// APIServer api server configuration
// 	APIServer = &configs.APIServerInfo{}
// )

// func getConf(dbSettings *configs.DbConnection, serverSetting *configs.APIServerInfo) {
// 	*dbSettings, *serverSetting, _ = configs.InitConfig()
// }

// func initAPI() {
// 	api.StartAPI(APIServer.Hostname, APIServer.Port, DbConnectionInfo)
// }

// func initDatastore() {
// 	user := DbConnectionInfo.User
// 	db := DbConnectionInfo.Database
// 	pass := DbConnectionInfo.Password
// 	host := DbConnectionInfo.Host
// 	port := DbConnectionInfo.Port
// 	datastores.Store().InitDatabase(user, db, pass, host, port)
// }

// func main() {
// 	getConf(DbConnectionInfo, APIServer)
// 	initDatastore()
// 	initAPI()
// }
