package config

import "embed"

//go:embed config.yml
var StandardConf embed.FS

//go:embed lang.json
var MCLangJson embed.FS
