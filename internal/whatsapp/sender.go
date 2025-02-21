package wapp

import (
	"bytes"
	"fmt"
	"net/http"

	BackupConfig "github.com/Crusher131/Bkp-Dados/internal/config"
)

func SendMsg(s bool, errmsg error) error {
	c := BackupConfig.GetConfig()
	url := c.Whatsapp.Server
	var subject string

	if s {
		subject = "*Backup cliente " + c.Backup.Client + " Finalizado com sucesso.*"
	} else {
		subject = fmt.Sprintf("*Backup Cliente %s erro:* %s", c.Backup.Client, errmsg.Error())
	}
	data := `{
		"query": "mutation partner_api_send_message{partner_api_send_message(recipient:\"` + c.Whatsapp.Number + `\" message:\"` + subject + `\" tipo:\"text\" sender_name:\"Monitoramento Hardtec\"){message}}",
		"variables": {},
		"operationName": "partner_api_send_message"
	}`

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("token", c.Whatsapp.ApiKeys)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println("Status Code:", resp.StatusCode)
	return nil
}
