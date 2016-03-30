package main

import (
	"flag"
	"log"
	"net/http"
)

//go:generate /bin/sh -c "cd ./root-fs && gopherjs build -v -o app.js"

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
)

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir("./root-fs")))
	err := http.ListenAndServe(*addrFlag, nil)
	if err != nil {
		log.Fatal(err)
	}
}
