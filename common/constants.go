package common

const (
	DatabaseSection  string = "database"
	DatabaseHost     string = "host"
	DatabaseUser     string = "user"
	DatabasePort     string = "port"
	DatabasePassword string = "password"
	DatabaseName     string = "name"
)

const (
	JwtSection string = "jwt"
	JwtSecret  string = "secret"
)

const (
	ServerSection string = "server"
	ServerListen  string = "listen"
)

type J map[string]interface{}
