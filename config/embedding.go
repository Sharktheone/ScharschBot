package config

import "embed"

//go:embed config.yml lang.json
var eFS embed.FS

func GetLang() ([]byte, error) {
	return eFS.ReadFile("lang.json")
}

func GetDefaultConf() ([]byte, error) {
	return eFS.ReadFile("config.yml")
}
