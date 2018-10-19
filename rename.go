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

	sftp, err := getSftpClient()
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
