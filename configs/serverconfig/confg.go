package serverconfig

type ServerConfiguration struct {
	PORT         string
	WriteTimeout int
	ReadTimeout  int
}

var serverConfig = ServerConfiguration{
	PORT:         "9000",
	WriteTimeout: 15,
	ReadTimeout:  15,
}

func GetConfig() ServerConfiguration {
	return serverConfig
}

func SetConfig(config ServerConfiguration) {
	serverConfig = config
}
