# go-web-plot-flot-embedded

A simple web server using the [flot](http://www.flotcharts.org/) library to plot.
Data is read from a `WebSocket` and pushed to the plot.
`flot` and `jquery` are embedded using `go-bindata-assetfs` and served locally.

Run like so:

```sh
$> cd 05-go-web-plot-flot-embedded && go get -v .
$> 05-go-web-plot-flot-embedded &
$> open http://localhost:5555
```
