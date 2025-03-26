package config

var Env = struct {
	SecretKey     string
	ServerAddress string
}{
	SecretKey:     "NOTSECRETKEY",
	ServerAddress: "SERVER_ADDRESS",
}
