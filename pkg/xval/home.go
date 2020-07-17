package xval

import (
	"fmt"
	"log"
	"net/http"
)

// Start 启动首页
func Start(port int) {
	mux := &http.ServeMux{}
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	log.Println("Home page at ", port)
	mux.HandleFunc("/", PageHandle)
	mux.HandleFunc("/app", AppHandle)
	log.Fatal(svr.ListenAndServe())
}
