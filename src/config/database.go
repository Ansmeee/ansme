package config

type databaseConfig struct {
	Driver   string
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func DBConfig() *databaseConfig {
	newConfig := new(databaseConfig)

	newConfig.Driver = Get("db_driver")
	newConfig.Host = Get("db_host")
	newConfig.Port = Get("db_port")
	newConfig.User = Get("db_user")
	newConfig.Pass = Get("db_pass")
	newConfig.Database = Get("database")

	return newConfig
}
