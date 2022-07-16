package conf

type Format struct {
	Reconfigure bool `yaml:"reconfigure"`
	Enabled     bool `yaml:"enabled"`
	Discord     struct {
		ServerID string
		Token    string
	}
	Whitelist struct {
		Enabled bool `yaml:"enabled"`
		Roles   struct {
			Enabled           bool   `yaml:"enabled"`
			RoleID            int64  `yaml:"roleID"`
			RemoveUserWithout bool   `yaml:"removeUserWithout"`
			RemoveTimeout     string `yaml:"removeTimeout"`
		}
		Pterodactyl struct {
			Enabled          bool   `yaml:"enabled"`
			APIKey           string `yaml:"APIKey"`
			PanelURL         string `yaml:"panelURL"`
			ServerID         string `yaml:"serverID"`
			WhitelistCommand string `yaml:"whitelistCommand"`
		}
		Luckperms struct {
			Enabled      bool   `yaml:"enabled"`
			SetRole      string `yaml:"setRole"`
			DatabaseHost string `yaml:"databaseHost"`
			DatabasePort uint16 `yaml:"databasePort"`
			DatabaseUser string `yaml:"databaseUser"`
			DatabasePass string `yaml:"databasePass"`
		}
	}
}
