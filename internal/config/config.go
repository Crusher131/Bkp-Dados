package BackupConfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Crusher131/logger"
	"gopkg.in/yaml.v3"
)

func Init(path, name string) error {
	c := defaultConfig()

	filepath := filepath.Join(path, name)

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		logger.Warn(fmt.Sprintf("arquivo: não existe %s", filepath))
		if err := createDefaultConfig(path, name, c); err != nil {
			return err
		}
	}
	if err := LoadConfig(filepath); err != nil {
		return err
	}

	return nil
}

func LoadConfig(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	if err = yaml.Unmarshal(data, &_cfg); err != nil {
		if syntarErr, ok := err.(*yaml.TypeError); ok {
			return fmt.Errorf("erro de sintaxe yaml:\n%s", syntarErr.Errors)
		}

	}
	return err
}

func createDefaultConfig(path, name string, c VConfig) (err error) {

	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	file, err := os.Create(path + "/" + name)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	if err = encoder.Encode(c); err != nil {
		return err
	}
	defer encoder.Close()

	return err
}

func defaultConfig() VConfig {
	_cfg := VConfig{
		Backup: struct {
			Client             string   `yaml:"client"`
			Retention          int      `yaml:"retention"`
			Backup_source      []string `yaml:"backup_source"`
			Backup_destination string   `yaml:"backup_destination"`
		}{
			Client:             "Nome Cliente",
			Retention:          7,
			Backup_source:      []string{"/dados", "/var/log"},
			Backup_destination: "/backup",
		},
		Whatsapp: struct {
			Server  string `yaml:"server"`
			ApiKeys string `yaml:"apikeys"`
			Send    bool   `yaml:"send"`
			Number  string `yaml:"number"`
		}{
			Server:  "",
			ApiKeys: "",
			Send:    false,
			Number:  "",
		},
		Configured: false,
	}
	return _cfg
}

func GetClient() string {
	return _cfg.Backup.Client
}
func GetRetention() int {
	return _cfg.Backup.Retention
}
func GetBackupDestination() string {

	return _cfg.Backup.Backup_destination
}
func GetBackupSource() []string {
	return _cfg.Backup.Backup_source
}
func GetConfig() VConfig {
	return *_cfg
}
func GetConfigured() bool {
	return _cfg.Configured
}
