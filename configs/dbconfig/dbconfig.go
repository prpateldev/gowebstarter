package dbconfig

type DatabaseConfig struct {
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBDbase string
}

var DBConfig DatabaseConfig = DatabaseConfig{
	DBHost:  "127.0.0.1",
	DBPort:  ":3306",
	DBUser:  "database user",
	DBPass:  "password for database",
	DBDbase: "database name",
}
