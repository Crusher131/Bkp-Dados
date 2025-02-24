package main

import (
	"fmt"

	"github.com/Crusher131/Bkp-Dados/internal/backup"
	vconf "github.com/Crusher131/Bkp-Dados/internal/config"
	"github.com/Crusher131/Bkp-Dados/internal/retention"
	wapp "github.com/Crusher131/Bkp-Dados/internal/whatsapp"
	"github.com/Crusher131/logger"
)

func main() {
	s := true
	conf_path := "./config"
	conf_file := "config.yaml"
	LogFile := "/home/jeferson/git/go/Bkp-Dados/backup.log"
	logger.Init(logger.SetLogFile(LogFile))
	logger.Info("Carregando Arquivos de configurações")
	if err := vconf.Init(conf_path, conf_file); err != nil {
		s = false
		wapp.SendMsg(s, err)
		logger.Fatal(err)
	}
	logger.Info("configurações carregadas com sucesso")
	if !vconf.GetConfigured() {
		logger.Fatal(fmt.Errorf("configure o arquivo config.yaml e altere o configured para true, para conseguir executar"))
	}
	logger.Info("iniciando backup")
	if err := backup.BackupCreate(); err != nil {
		s = false
		wapp.SendMsg(s, err)
		logger.Error(fmt.Errorf("erro ao fazer o backup dos arquivos"))
		logger.Fatal(fmt.Errorf("%e", err))
	}
	logger.Info("backup Efetuado com sucesso!")
	logger.Info("removendo arquivos com data maior que a retenção")
	if err := retention.RetentionRemove(); err != nil {
		s = false
		wapp.SendMsg(s, err)
		logger.Error(fmt.Errorf("erro ao remover arquivos antigos"))
		logger.Warn(fmt.Sprint(err))
	}

	wapp.SendMsg(s, nil)

	/* ##TODO
	Gerar/Gerenciar logs															DONE
	Gerar/Gerenciar config.yaml utilizar *viper*?									DONE
	Definir pastas/arquivos a serem feitos backup(Pode ser um array de objetos)		DONE
	Gerar Arquivo ZIP																DONE
	Verificar Integridade															DONE
	Remover backups utilizando tempo limite configurado no config.yaml				DONE
	Nome do arquivo de backup Utilizando nome do cliente e data						DONE
	Enviar notificação via Whatsapp													DONE
	*/

}
