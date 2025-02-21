package retention

import (
	"fmt"
	"os"
	"time"

	"github.com/Crusher131/Bkp-Dados/internal/cmd"
	BackupConfig "github.com/Crusher131/Bkp-Dados/internal/config"
	"github.com/Crusher131/logger"
)

func RetentionRemove() error {
	c := BackupConfig.GetConfig()
	os.Chdir(c.Backup.Backup_destination)
	files, err := getFiles(c.Backup.Backup_destination)
	if err != nil {
		return err
	}
	for _, v := range files {
		duration, err := getInfoFiles(v)
		if err != nil {
			return err
		}

		timeRetention := float64(c.Backup.Retention * 24)
		if duration.Hours() > timeRetention {
			logger.Info(fmt.Sprintf("Arquivo %s tem a duração maior que o tempo de retenção: %.2f", v, duration.Hours()))
			logger.Info(fmt.Sprintf("Tempo de retenção: %.2f\tTempo do arquivo: %.2f", timeRetention, duration.Hours()))
			logger.Info("O arquivo será removido")
			removeFiles(v)
		}

	}
	return nil

}
func removeFiles(filename string) error {
	err := cmd.CmdExec("rm", filename, "-rfv")
	return err

}

func getFiles(folderpath string) (filesnames []string, err error) {

	files, err := os.ReadDir(folderpath)
	if err != nil {
		return nil, err
	}
	var fileName []string

	for _, entry := range files {
		if !entry.IsDir() {
			fileName = append(fileName, entry.Name())
		}
	}
	return fileName, err
}

func getInfoFiles(filename string) (duration time.Duration, err error) {
	info, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	modtime := info.ModTime()
	duration = time.Since(modtime)

	return
}
