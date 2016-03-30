# go-web-plot-flot-static

A simple web server using the [flot](http://www.flotcharts.org/) library to plot.
Data is read from a `WebSocket` and pushed to the plot.
`flot` and `jquery` are served locally.

Run like so:

```sh
$> cd 04-go-web-plot-flot-static; go run ./main.go
$> open http://localhost:5555
```
