package main

import (
	"flag"
	"log"
	"math"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
	datac    = make(chan Data)
)

type Data struct {
	X   float64 `json:"x"`
	Sin float64 `json:"sin"`
	Cos float64 `json:"cos"`
}

func main() {
	flag.Parse()

	done := make(chan bool)
	go generate(datac, done)

	http.Handle("/", http.FileServer(http.Dir("./root-fs")))
	http.Handle("/data", websocket.Handler(dataHandler))
	err := http.ListenAndServe(*addrFlag, nil)
	if err != nil {
		done <- true
		log.Fatal(err)
	}
}

func generate(datac chan Data, done chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	x := -0.5
	for {
		select {
		case <-ticker.C:
			x += 0.5
			datac <- Data{x, math.Sin(x), math.Cos(x)}
		case <-done:
			return
		}
	}
}

func dataHandler(ws *websocket.Conn) {
	for data := range datac {
		err := websocket.JSON.Send(ws, data)
		if err != nil {
			log.Printf("error sending data: %v\n", err)
			return
		}
	}
}
