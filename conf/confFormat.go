package conf

type Format struct {
	Reconfigure bool `yaml:"reconfigure"`
	Enabled     bool `yaml:"enabled"`
	Discord     struct {
		ServerID              string `yaml:"serverID"`
		Token                 string `yaml:"token"`
		WhitelistRemoveRoleID string `yaml:"adminWhitelistRemoveRoleID"`
		WhitelistWhoisRoleID  string `yaml:"adminWhitelistWhoisRoleID"`
	}
	Pterodactyl struct {
		Enabled          bool   `yaml:"enabled"`
		APIKey           string `yaml:"APIKey"`
		PanelURL         string `yaml:"panelURL"`
		ServerID         string `yaml:"serverID"`
		WhitelistCommand string `yaml:"whitelistCommand"`
	}
	Whitelist struct {
		Enabled bool `yaml:"enabled"`
		Mongodb struct {
			MongodbHost           string `yaml:"mongodbHost"`
			MongodbPort           uint16 `yaml:"mongodbPort"`
			MongodbUser           string `yaml:"mongodbUser"`
			MongodbPass           string `yaml:"mongodbPass"`
			MongodbDatabaseName   string `yaml:"mongodbDatabaseName"`
			MongodbCollectionName string `yaml:"mongodbCollectionName"`
		}

		Roles struct {
			Enabled           bool   `yaml:"enabled"`
			RoleID            int64  `yaml:"roleID"`
			RemoveUserWithout bool   `yaml:"removeUserWithout"`
			RemoveTimeout     string `yaml:"removeTimeout"`
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
