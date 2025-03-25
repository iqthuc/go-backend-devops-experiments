package config

var Env = struct {
	SecretKey     string
	ServerAddress string
}{
	SecretKey:     "SECRETKEY",
	ServerAddress: "SERVER_ADDRESS",
}
