package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
)

func main() {
	flag.Parse()
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(*addrFlag, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world 2.0\n")
}
