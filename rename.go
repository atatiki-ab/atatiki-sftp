package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	routes = append(routes, Route{"renameHandler", "POST", "/rename", renameHandler})
}

func renameHandler(w http.ResponseWriter, r *http.Request) {

	command := struct {
		OldName string `json:"oldname"`
		NewName string `json:"newname"`
	}{}
	if r.Body == nil {
		logger.Info("No body")
		http.Error(w, "No body", 500)
		return
	}
	logger.Info("Got the data")
	if err := json.NewDecoder(r.Body).Decode(&command); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	logger.Info("Could decode the data")

	sftp, err := getSftpClient()
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer sftp.Close()
	logger.Info("Could get a sftp client connection")

	//pwd, _ := sftp.Getwd()
	//logger.Info(pwd)
	if err := sftp.Rename(command.OldName, command.NewName); err != nil {
		logger.Info("Got an error from the remote server")
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Done")

}
