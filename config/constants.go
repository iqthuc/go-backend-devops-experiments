package config

var Env = struct {
	SecretKey     string
	ServerAddress string
	RedisAddr     string
	RedisPassword string
}{
	SecretKey:     "NOTSECRETKEY",
	ServerAddress: "SERVER_ADDRESS",
	RedisAddr:     "REDIS_ADDRESS",
	RedisPassword: "REDIS_PASSWORD",
}
