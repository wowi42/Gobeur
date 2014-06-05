package main

import (
	"io"
	"net/http"
	"log"
	"fmt"
)
func WebDAVServer(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(4096)
	fmt.Println(req.MultipartForm)
	for _, fileHeaders := range req.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			file, _ := fileHeader.Open()
			path := fmt.Sprintf("files/%s", fileHeader.Filename)
			fmt.Println(file)
			fmt.Println(path)
		}
	}
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/webdav", WebDAVServer)
	err := http.ListenAndServe(":23456", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
