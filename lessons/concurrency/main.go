package main

import (
	"fmt"
	"log"
	"net/http"
)

func serverApp()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello World")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serverDebug()  {
	if err := http.ListenAndServe("0.0.0.0:8081", http.DefaultServeMux); err != nil {
		// log.Fatal 调用了os.Exit 会无条件终止程序；defers不会被调用到
		log.Fatal(err)
	}
}

func main() {
	go serverDebug()
	serverApp()
	select {}
}
