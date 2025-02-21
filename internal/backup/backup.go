package backup

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/Crusher131/Bkp-Dados/internal/cmd"
	BackupConfig "github.com/Crusher131/Bkp-Dados/internal/config"
)

func BackupCreate() error {
	c := BackupConfig.GetConfig()

	time := time.Now()
	year, month, day := time.Date()

	backupname := fmt.Sprintf("%s-%d-%d-%d.tar.gz", c.Backup.Client, day, month, year)

	Dest := filepath.Join(c.Backup.Backup_destination, backupname)

	args := []string{"-czf", Dest}
	args = append(args, c.Backup.Backup_source...)

	if err := cmd.CmdExec("tar", args...); err != nil {

		return err

	}

	return nil
}
