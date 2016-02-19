package main

import (
	"flag"
	"fmt"
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

	http.HandleFunc("/", plotHandle)
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

func plotHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
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

const page = `
<html>
	<head>
		<title>Plotting stuff with Flot</title>
		<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
		<script src="//cdnjs.cloudflare.com/ajax/libs/flot/0.8.3/jquery.flot.min.js"></script>
		<script type="text/javascript">
		var sock = null;
		var sinplot = {
			label: "sin(x)",
			data: [],
		};
		var cosplot = {
			label: "cos(x)",
			data: [],
		};

		function update() {
			var p1 = $.plot("#my-sin-plot", [sinplot]);
			p1.setupGrid(); // needed as x-axis changes
			p1.draw();

			var cos = $.plot("#my-cos-plot", [cosplot]);
			cos.setupGrid();
			cos.draw();
		};

		window.onload = function() {
			sock = new WebSocket("ws://localhost:5555/data");

			sock.onmessage = function(event) {
				var data = JSON.parse(event.data);
				console.log("data: "+JSON.stringify(data));
				sinplot.data.push([data.x, data.sin]);
				cosplot.data.push([data.x, data.cos]);
				update();
			};
		};

		</script>

		<style>
		.my-plot-style {
			width: 400px;
			height: 200px;
			font-size: 14px;
			line-height: 1.2em;
		}
		</style>
	</head>

	<body>
		<div id="header">
			<h2>My plot</h2>
		</div>

		<div id="content">
			<div id="my-sin-plot" class="my-plot-style"></div>
			<br>
			<div id="my-cos-plot" class="my-plot-style"></div>
		</div>
	</body>
</html>
`
