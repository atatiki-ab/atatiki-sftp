package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func init() {
	routes = append(routes, Route{"renameHandler", "POST", "/rename", renameHandler})
}

func renameHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//logger.Info("Here")
	//fmt.Println(settings)

	command := struct {
		OldName string `json:"oldname"`
		NewName string `json:"newname"`
	}{}
	if r.Body == nil {
		logger.Info("No body")
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		logger.Error(err)
		return
	}

	sshConfig := &ssh.ClientConfig{
		User: settings.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(settings.Password),
		},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	host := "localhost:22"
	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		logger.Error(err)
		return
	}

	sftp, err := sftp.NewClient(client)
	if err != nil {
		logger.Error(err)
		return
	}
	defer sftp.Close()

	pwd, _ := sftp.Getwd()
	logger.Info(pwd)
	if err := sftp.Rename(command.OldName, command.NewName); err != nil {
		logger.Error(err)
		fmt.Fprintf(w, "Error")
		return
	}
	fmt.Fprintf(w, "Done")

}
