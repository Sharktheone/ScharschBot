package conf

type Format struct {
	Reconfigure bool `yaml:"reconfigure"`
	Enabled     bool `yaml:"enabled"`
	Discord     struct {
		ServerID              string `yaml:"serverID"`
		Token                 string `yaml:"token"`
		WhitelistRemoveRoleID string `yaml:"adminWhitelistRemoveRoleID"`
		WhitelistWhoisRoleID  string `yaml:"adminWhitelistWhoisRoleID"`
		WhitelistBanRoleID    string `yaml:"adminWhitelistBanRoleID"`
		WhitelistServerRoleID string `yaml:"whitelistServerRoleID"`
		EmbedErrorIcon        string `yaml:"embedErrorIcon"`
		EmbedErrorAuthorURL   string `yaml:"embedErrorAuthorURL"`
	}
	Pterodactyl struct {
		Enabled          bool   `yaml:"enabled"`
		APIKey           string `yaml:"APIKey"`
		PanelURL         string `yaml:"panelURL"`
		ServerID         string `yaml:"serverID"`
		WhitelistCommand string `yaml:"whitelistCommand"`
	}
	Whitelist struct {
		Enabled     bool `yaml:"enabled"`
		MaxAccounts int  `yaml:"maxAccounts"`
		Mongodb     struct {
			MongodbHost                    string `yaml:"mongodbHost"`
			MongodbPort                    uint16 `yaml:"mongodbPort"`
			MongodbUser                    string `yaml:"mongodbUser"`
			MongodbPass                    string `yaml:"mongodbPass"`
			MongodbDatabaseName            string `yaml:"mongodbDatabaseName"`
			MongodbWhitelistCollectionName string `yaml:"mongodbWhitelistCollectionName"`
			MongodbBanCollectionName       string `yaml:"mongodbBanCollectionName"`
		}

		Roles struct {
			Enabled            bool   `yaml:"enabled"`
			ServerRoleID       string `yaml:"serverRoleID"`
			RemoveRoleOthersID string `yaml:"removeRoleOthersID"`
			ListUserRoleID     string `yaml:"listUserRoleID"`
			WhoisRoleID        string `yaml:"whoisRoleID"`
			RemoveAllRoleID    string `yaml:"removeAllRoleID"`
			RemoveUserWithout  bool   `yaml:"removeUserWithout"`
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
