package main

import (
	"fmt"
	"net/http"
)

func init() {
	routes = append(routes, Route{"testHandler", "GET", "/pwd", testHandler})
}

func testHandler(w http.ResponseWriter, r *http.Request) {

	sftp, err := getSftpClient()
	if err != nil {
		logger.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}
	defer sftp.Close()

	pwd, _ := sftp.Getwd()
	fmt.Fprintf(w, pwd)

}
