package common

const (
	DatabaseSection  string = "database"
	DatabaseDialect  string = "dialect"
	DatabasePath     string = "path"
	DatabaseHost     string = "host"
	DatabaseUser     string = "user"
	DatabasePort     string = "port"
	DatabasePassword string = "password"
	DatabaseName     string = "name"
)

const (
	JwtSection    string = "jwt"
	JwtSecret     string = "secret"
	JwtSignMethod string = "crypto"
)

const (
	ServerSection string = "server"
	ServerListen  string = "listen"
)

type JSON map[string]interface{}
