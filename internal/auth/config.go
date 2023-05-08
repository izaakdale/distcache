package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile         = configFile("ca-creds/tls.crt")
	ServerCertFile = configFile("server-creds/tls.crt")
	ServerKeyFile  = configFile("server-creds/tls.key")
	ClientCertFile = configFile("client-creds/tls.crt")
	ClientKeyFile  = configFile("client-creds/tls.key")
)

func init() {
	if os.Getenv("ENV") == "local" {
		CAFile = configFile("ca.pem")
		ServerCertFile = configFile("server.pem")
		ServerKeyFile = configFile("server-key.pem")
		ClientCertFile = configFile("client.pem")
		ClientKeyFile = configFile("client-key.pem")
	}
}

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homedir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homedir, ".distcache", filename)
}
