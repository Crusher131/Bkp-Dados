package BackupConfig

type VConfig struct {
	Backup     Backup
	Whatsapp   Whatsapp
	Configured bool `yaml:"configured"`
}

type Backup struct {
	Client             string   `yaml:"client"`
	Retention          int      `yaml:"retention"`
	Backup_source      []string `yaml:"backup_source"`
	Backup_destination string   `yaml:"backup_destination"`
}

type Whatsapp struct {
	Server  string `yaml:"server"`
	ApiKeys string `yaml:"apikeys"`
	Send    bool   `yaml:"send"`
	Number  string `yaml:"number"`
}

var _cfg *VConfig
