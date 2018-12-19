package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func init() {
	routes = append(routes, Route{"downloadHandler", "POST", "/download", downloadHandler})
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		logger.Info("No body")
		http.Error(w, "No body", 500)
		return
	}
	f := struct {
		PathAndFileName          string `json:"pathAndFileName"`
		RemoveWindowsLineEndings bool   `json:"removeWindowsLineEndings"`
		ConvertFrom8859          bool   `json:"convertFrom8859"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	sftp, err := getSftpClient()
	if err != nil {
		logger.Error(err)
		return
	}
	defer sftp.Close()

	file, err := sftp.Open(f.PathAndFileName)
	if err != nil {
		logger.Error(err)
		return
	}
	buff := &bytes.Buffer{}
	if _, err := file.WriteTo(buff); err != nil {
		logger.Error(err)
		return
	}
	str := buff.String()
	if f.ConvertFrom8859 {
		str, err = charmap.ISO8859_1.NewDecoder().String(str)
	}

	if f.RemoveWindowsLineEndings {
		str = strings.Replace(str, "\r\n", "\n", -1)
	}
	fmt.Fprintf(w, str)
}
