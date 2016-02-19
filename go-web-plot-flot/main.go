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
	http.HandleFunc("/", plotHandle)
	err := http.ListenAndServe(*addrFlag, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func plotHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, page)
}

const page = `
<html>
	<head>
		<title>Plotting stuff with Flot</title>
		<script src="//cdnjs.cloudflare.com/ajax/libs/jquery/2.0.3/jquery.min.js"></script>
		<script src="//cdnjs.cloudflare.com/ajax/libs/flot/0.8.2/jquery.flot.min.js"></script>
		<script type="text/javascript">
		$(function() {
			var plot = {
				label: "sin(x)",
				data: [],
				clickable: true,
				hoverable: true,
			};
			var options = {
				grid: {
					hoverable: true,
				},
			};

			for (var i = 0; i < 20; i += 0.5) {
				plot.data.push([i, Math.sin(i)]);
			}

			$.plot("#my-plot", [plot], options);
		
		});

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
			<div id="my-plot" class="my-plot-style"></div>
		</div>
	</body>
</html>
`
