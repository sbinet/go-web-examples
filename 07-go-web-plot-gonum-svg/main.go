package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgsvg"
)

var (
	addrFlag = flag.String("addr", ":5555", "server address:port")
	datac    = make(chan plots)
)

type plots struct {
	Sin string `json:"sin"`
	Cos string `json:"cos"`
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

func generate(datac chan plots, done chan bool) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	table := make([]plotter.XYs, 2)
	dx := 0.05
	x := -dx
	for {
		select {
		case <-ticker.C:
			x += dx
			table[0] = append(table[0], struct{ X, Y float64 }{x, math.Sin(x)})
			table[1] = append(table[1], struct{ X, Y float64 }{x, math.Cos(x)})

			sinSVG := plotSin(table[0])
			cosSVG := plotCos(table[1])

			datac <- plots{sinSVG, cosSVG}
		case <-done:
			return
		}
	}
}

func plotSin(data plotter.XYs) string {

	sin, err := plot.New()
	if err != nil {
		panic(err)
	}
	sin.X.Label.Text = "Angle (radians)"
	sin.Y.Label.Text = "Sin(x)"

	line, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Color = color.RGBA{255, 0, 0, 255}
	line.LineStyle.Width = vg.Points(1)

	sin.Add(line)
	sin.Add(plotter.NewGrid())

	return renderSVG(sin)
}

func plotCos(data plotter.XYs) string {
	cos, err := plot.New()
	if err != nil {
		panic(err)
	}
	cos.X.Label.Text = "Angle (radians)"
	cos.Y.Label.Text = "Cos(x)"

	line, err := plotter.NewLine(data)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Color = color.RGBA{0, 0, 255, 255}
	line.LineStyle.Width = vg.Points(1)

	cos.Add(line)
	cos.Add(plotter.NewGrid())

	return renderSVG(cos)
}

func renderSVG(p *plot.Plot) string {
	size := 10 * vg.Centimeter
	canvas := vgsvg.New(size, size/vg.Length(math.Phi))
	p.Draw(draw.New(canvas))
	out := new(bytes.Buffer)
	_, err := canvas.WriteTo(out)
	if err != nil {
		panic(err)
	}
	return string(out.Bytes())
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
		<title>Plotting stuff with gonum/plot</title>
		<script type="text/javascript">
		var sock = null;
		var sinplot = "";
		var cosplot = "";

		function update() {
			var p1 = document.getElementById("my-sin-plot");
			p1.innerHTML = sinplot;

			var p2 = document.getElementById("my-cos-plot");
			p2.innerHTML = cosplot;
		};

		window.onload = function() {
			sock = new WebSocket("ws://"+location.host+"/data");

			sock.onmessage = function(event) {
				var data = JSON.parse(event.data);
				//console.log("data: "+JSON.stringify(data));
				sinplot = data.sin;
				cosplot = data.cos;
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
