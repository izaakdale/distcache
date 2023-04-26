package config

import (
	"os"
	"path/filepath"
)

var (
	CAFile             = configFile("ca-creds/tls.crt")
	ServerCertFile     = configFile("server-creds/tls.crt")
	ServerKeyFile      = configFile("server-creds/tls.key")
	RootClientCertFile = configFile("client-creds/tls.crt")
	RootClientKeyFile  = configFile("client-creds/tls.key")
)

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
