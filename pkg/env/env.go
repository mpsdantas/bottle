package env

import (
	"os"

	"github.com/mpsdantas/bottle/pkg/definition"
)

func init() {
	Bottle := definition.Bottle()

	Environment = Get("APPLICATION_ENV", Local)
	Scope = Get("APPLICATION_SCOPE", Server)
	Port = Get("PORT", "8080")
	Version = Get("APPLICATION_VERSION", "develop")
	Application = Bottle["application"]
}

var (
	Application string
	Environment string
	Scope       string
	Port        string
	Version     string
)

const (
	Local string = "local"
	Stage string = "stage"
	Prod  string = "prod"
)

const (
	Server     string = "server"
	Subscriber string = "subscriber"
	Cronjob    string = "cronjob"
)

func Get(key string, defaults ...string) string {
	val := os.Getenv(key)
	if val == "" && len(defaults) > 0 {
		return defaults[0]
	}

	return val
}
