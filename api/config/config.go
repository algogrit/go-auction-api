package config

import (
	"os"
	"strconv"
)

//AppConfig Application Configuration Structure
type AppConfig struct {
	dbhost         string
	dbname         string
	dbport         int
	connectionPool int
	appSecret      string
	appServerPort  int
}

//GetAppConfig Return pointer to APPConfig for DEV/PROD
func GetAppConfig() *AppConfig {
	appEnv := os.Getenv("APP_ENV")
	switch appEnv {
	case "PROD":
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT_PROD"))
		conPool, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_POOL_PROD"))
		serverPort, _ := strconv.Atoi(os.Getenv("APP_SERVER_PORT_PROD"))

		return &AppConfig{
			dbhost:         os.Getenv("DB_HOST_PROD"),
			dbname:         os.Getenv("DB_NAME_PROD"),
			dbport:         dbPort,
			connectionPool: conPool,
			appSecret:      os.Getenv("APP_SECRET_PROD"),
			appServerPort:  serverPort,
		}
	default:
		dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT_DEV"))
		conPool, _ := strconv.Atoi(os.Getenv("DB_CONNECTION_POOL_DEV"))
		serverPort, _ := strconv.Atoi(os.Getenv("APP_SERVER_PORT_DEV"))

		return &AppConfig{
			dbhost:         os.Getenv("DB_HOST_DEV"),
			dbname:         os.Getenv("DB_NAME_DEV"),
			dbport:         dbPort,
			connectionPool: conPool,
			appSecret:      os.Getenv("APP_SECRET_DEV"),
			appServerPort:  serverPort,
		}

	}
}

func (cfg *AppConfig) GetDatabaseHostname() string {
	return cfg.dbhost
}
func (cfg *AppConfig) GetDatabaseName() string {
	return cfg.dbname
}
func (cfg *AppConfig) GetAppSecret() string {
	return cfg.appSecret
}
func (cfg *AppConfig) GetDatabasePort() string {
	return strconv.Itoa(cfg.dbport)
}
func (cfg *AppConfig) GetConnectionPool() int {
	return cfg.connectionPool
}
func (cfg *AppConfig) GetAppServerPort() string {
	return strconv.Itoa(cfg.appServerPort)
}
