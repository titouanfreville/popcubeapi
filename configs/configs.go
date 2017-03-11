package configs

import (
	"log"
	"os"
)

// DbConnection information to connect to DB
type DbConnection struct {
	User     string
	Database string
	Password string
	Host     string
	Port     string
}

// APIServerInfo information on API server
type APIServerInfo struct {
	Hostname string
	Port     string
}

// InitConfig get configuration for project
func InitConfig() (DbConnection, APIServerInfo, string) {
	// Default configurations
	dbConnection := DbConnection{
		User:     "root",
		Database: "popcube_test",
		Password: "popcube_dev",
		Host:     "0.0.0.0",
		Port:     "3306",
	}
	APIServer := APIServerInfo{
		Hostname: "",
		Port:     "3000",
	}
	// Dev secret
	secret := "AAAAB3NzaC1yc2EAAAADAQABAAACAQCtdGt4uK8e1CEcTVZXSRJ9pRHxdeYBxq4oTh20DKH7exoikkEPbSAn34ZJPVVRdPMndg8Qg5xxHnwAtYvzYbxNWAxqYqvvvCKLJtjTS2dMeNLVz3FYD80MSJX3Tr5gpK7hHq9EEWB99onqMKDHlF3ZM3dBjwZH3mP7sWlqcdKc6lP9MGPsrpnXmBx3C4CSB7muMl8hF+4263gtS1oXHT0E16NFP3IgBNmvYavmOYSlqHs9NU7lZtNVbLbIZ2SCVrOJlcSKddvaMzIhXgRIK58VzbsqqaeVBTMrxrJopjLha2aTSe9luxOJZCf1foQKVf7eWPp4FK/zSSDMJbSX6+vsE1jFbuFF2dYmf8QW1UdDslZtQuCLzB4rqBmOiFx77DIyuZMMt5bjTi02nPYZL5Fo4vupcoV552QC6jyUG3nAoY28yPGmhKBb0EpbCd/qiroIAs5mXhaPGZriqq8DDRbqstHkubfXjDkZ6vWRDnCUfSioMky/bEC1X2KaMt/E0tpw8aWiIZXAble+CIWfo2HUj2GE/Y3Gf8f/A14Ec2E+Uz4xARcTL4UfopNU2P3Bxhz/KoIZFXYacKBphATsp+HB6sMKF5HJ+tn6mS0JFdgIpcClVMliap4zz6M92FOyyRW0wBHua6gOI+5nEMS2BDLBwTmw5otXOTFV8DaFNQzaiQ"
	// Default host for DB in Docker containers
	if os.Getenv("ENVTYPE") == "container" {
		log.Print("<><><><> Setting host to container default \n")
		dbConnection.Host = "database"
	}

	if newSecret := os.Getenv("POPCUBESECRET"); newSecret != "" {
		secret = newSecret
	}

	// Get values set in env
	if apiPort := os.Getenv("API_PORT"); apiPort != "" {
		log.Print("<><><><> Setting api port \n")
		APIServer.Port = apiPort
	}
	if apiHostname := os.Getenv("API_HOST"); apiHostname != "" {
		log.Print("<><><><> Setting api hostname \n")
		APIServer.Hostname = apiHostname
	}
	// Will be erased if user is not root
	if dbRootPassword := os.Getenv("MYSQL_ROOT_PASSWORD"); dbRootPassword != "" {
		log.Print("<><><><> Setting db root password \n")
		dbConnection.Password = dbRootPassword
	}
	if dbUser := os.Getenv("MYSQL_USER"); dbUser != "" {
		log.Print("<><><><> Setting db user and user password \n")
		dbConnection.User = dbUser
		// Can be empty. Should be define when user is define
		dbConnection.Password = os.Getenv("MYSQL_PASSWORD")
	}
	if dbName := os.Getenv("MYSQL_DATABASE"); dbName != "" {
		log.Print("<><><><> Setting db name \n")
		dbConnection.Database = dbName
	}
	if dbPort := os.Getenv("MYSQL_PORT"); dbPort != "" {
		log.Print("<><><><> Setting db port \n")
		dbConnection.Port = dbPort
	}
	if dbHost := os.Getenv("MYSQL_HOST"); dbHost != "" {
		log.Print("<><><><> Setting db host \n")
		dbConnection.Host = dbHost
	}

	// Return new configs
	return dbConnection, APIServer, secret
}
