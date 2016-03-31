# go-web-plot-gonum-svg

A simple web server using the [gonum/plot](https://github.com/gonum/plot) library to plot.
The plot is created on the `Go` side (_i.e._ server-side), plotted as `SVG` and pushed to the client over a `WebSocket`.

Run like so:

```sh
$> 07-go-web-plot-gonum-svg &
$> open http://localhost:5555
```
