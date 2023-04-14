package flags

import (
	"flag"
)

func init() {
	flag.Bool("pprof", false, "enable pprof")
	flag.Int("pprof-port", 6060, "port for pprof")
	flag.Int("apiPort", 8080, "Port for the API")
	flag.String("configPath", "config.yml", "Path to the config file (default: config/config.yml)")
	flag.String("guild", "", "Guild ID")
	flag.String("token", "", "Discord Bot Token")
	flag.String("minecraftLangPath", "internal", "Path to lang.json")
	flag.String("apiUser", "", "Username for the API")
	flag.String("apiPassword", "", "Password for the API")

	flag.Parse()
}

func String(name string) *string {
	v := flag.Lookup(name).Value.String()
	return &v
}
func Int(name string) *int {
	v := flag.Lookup(name).Value.(flag.Getter).Get().(int)
	return &v
}
func Bool(name string) *bool {
	v := flag.Lookup(name).Value.(flag.Getter).Get().(bool)
	return &v
}

func StringWithFallback(name string, fallback *string) *string {
	v := String(name)
	if *v == "" {
		return fallback
	}
	return v
}

func IntWithFallback(name string, fallback *int) *int {
	v := Int(name)
	if *v == 0 {
		return fallback
	}
	return v
}
