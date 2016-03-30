// +build js

package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/sbinet/go-web-examples/06-go-web-gopherjs/pkg"
)

func main() {
	doc := js.Global.Get("document")
	doc.Get("body").Set("innerHTML", "GopherJS demo starting... (content will appear in 2s)")
	<-time.After(2 * time.Second)
	str := "you"
	doc.Get("body").Set("innerHTML", pkg.Hello(str)+"\n<p>An <em>alert</em> should pop up in 5s</p>")
	<-time.After(5 * time.Second)
	js.Global.Call("alert", pkg.Hello(str))
}
