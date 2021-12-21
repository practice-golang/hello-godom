package main // import "server"

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/NYTimes/gziphandler"
)

//go:embed dist/*
var content embed.FS

func main() {
	listen := "127.0.0.1:5525"

	fsys, err := fs.Sub(content, "dist")
	if err != nil {
		panic(err)
	}

	println("Server listening: " + listen)

	// http.ListenAndServe(listen, http.FileServer(http.FS(fsys)))
	http.ListenAndServe(listen, gziphandler.GzipHandler(http.FileServer(http.FS(fsys))))
}
