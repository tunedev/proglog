package config

import (
	"os"
	"path/filepath"
)

var (
	CAFILE = configFile("ca.pem")
	ServerCertFile = configFile("server.pem")
	ServerKeyFile = configFile("server-key.pem")
	RootClientCertFile = configFile("root-client.pem")
	NobodyClientCertFile = configFile("nobody-client.pem")
	RootClientKeyFile = configFile("root-client-key.pem")
	NobodyClientKeyFile = configFile("nobody-client-key.pem")
	ACLPolicyFile = configFile("policy.csv")
	ACLModelFile = configFile("model.conf")
)

func configFile(filename string) string {
	if dir := os.Getenv("CONFIG_DIR"); dir != "" {
		return filepath.Join(dir, filename)
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return filepath.Join(homeDir, ".proglog", filename)
}
