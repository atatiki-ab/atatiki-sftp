package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type remoteFile struct {
	PathAndFilename       string `json:"pathAndFilename"`
	Content               string `json:"content"`
	AddWindowsLineEndings bool   `json:"addWindowsLineEndings"`
	ConvertTo8859         bool   `json:"convertTo8859"`
}

func init() {
	routes = append(routes, Route{"copyRemoteHandler", "POST", "/copyremote", copyRemoteHandler})
}

func copyRemoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		logger.Info("No body")
		http.Error(w, "No body", 500)
		return
	}
	rf := remoteFile{}
	if err := json.NewDecoder(r.Body).Decode(&rf); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	if err := copyRemote(rf); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Fprintf(w, "Done")
}

func copyRemote(rf remoteFile) error {
	sftp, err := getSftpClient()
	if err != nil {
		return err
	}
	defer sftp.Close()

	file, err := sftp.Create(rf.PathAndFilename)
	if err != nil {
		return err
	}
	defer file.Close()

	if rf.AddWindowsLineEndings {
		rf.Content = strings.Replace(rf.Content, "\n", "\r\n", -1)
	}

	var b []byte
	if rf.ConvertTo8859 {
		if b, err = charmap.ISO8859_1.NewEncoder().Bytes([]byte(rf.Content)); err != nil {
			return err
		}
	} else {
		b = []byte(rf.Content)
	}

	if _, err := file.Write(b); err != nil {
		return err
	}
	return nil
}
