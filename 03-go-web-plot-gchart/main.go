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
		<title>Plotting stuff with Google-Charts</title>
		<script type="text/javascript" src="https://www.gstatic.com/charts/loader.js"></script>
		<script type="text/javascript">
			google.charts.load('current', {packages: ['corechart', 'line']});
			google.charts.setOnLoadCallback(initDrawSin);
			google.charts.setOnLoadCallback(initDrawCos);

			var sindata = null;
			var sinplot = null;
			function initDrawSin() {
				sindata = new google.visualization.DataTable();
				sindata.addColumn("number", "time");
				sindata.addColumn("number", "sine");

				sinplot = new google.charts.Line(document.getElementById("my-sin-plot"));
				drawSin();
			};

			var cosdata = null;
			var cosplot = null;
			function initDrawCos() {
				cosdata = new google.visualization.DataTable();
				cosdata.addColumn("number", "time");
				cosdata.addColumn("number", "cosine");

				cosplot = new google.charts.Line(document.getElementById("my-cos-plot"));
				drawCos();
			};

			function drawSin() {
				sinplot.draw(sindata, {
					hAxis: {
						title: "Time",
					},
					vAxis: {
						title: "Sine",
					},
					legend: {
						position: "none",
					},
					chart: {
						title: "Sine",
					},
				});
			};

			function drawCos() {
				cosplot.draw(cosdata, {
					hAxis: {
						title: "Time",
					},
					vAxis: {
						title: "Cosine",
					},
					legend: {
						position: "none",
					},
					chart: {
						title: "Cosine",
					},
				});
			};


			var sock = null;

			function update() {
				drawSin();
				drawCos();
			};

			window.onload = function() {
				sock = new WebSocket("ws://"+location.host+"/data");

				sock.onmessage = function(event) {
					var data = JSON.parse(event.data);
					console.log("data: "+JSON.stringify(data));
					sindata.addRows([[data.x, data.sin]]);
					cosdata.addRows([[data.x, data.cos]]);
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
