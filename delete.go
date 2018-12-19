package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func init() {
	routes = append(routes, Route{"deleteHandler", "POST", "/delete", deleteHandler})
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		logger.Info("No body")
		http.Error(w, "No body", 500)
		return
	}
	f := struct {
		PathAndFileName string `json:"pathAndFileName"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	if err := deleteRemote(f.PathAndFileName); err != nil {
		logger.Info(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Done")
}

func deleteRemote(pathAndFileName string) error {
	sftp, err := getSftpClient()
	if err != nil {
		return err
	}
	defer sftp.Close()

	return sftp.Remove(pathAndFileName)
}
