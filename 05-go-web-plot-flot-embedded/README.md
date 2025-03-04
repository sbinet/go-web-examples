# go-web-plot-flot-embedded

A simple web server using the [flot](http://www.flotcharts.org/) library to plot.
Data is read from a `WebSocket` and pushed to the plot.
`flot` and `jquery` are embedded using `go-bindata-assetfs` and served locally.

Run like so:

```sh
$> go get -v codeberg.org/sbinet/go-web-examples/05-go-web-plot-flot-embedded
$> cd /anywhere && 05-go-web-plot-flot-embedded &
$> open http://localhost:5555
```

`05-go-web-plot-flot-embedded` is exactly like `04-go-web-plot-flot-static` except it serves the content from an embedded set of resources (instead of serving them directly from the filesystem.)
The embedded resources have been created via [go-bindata-assetfs](https://github.com/elazarl/go-bindata-assetfs) and served via an `assetFS()` call which returns a value implementing the [net/http.Filesystem](https://godoc.org/net/http#FileSystem) interface.
