//go:build js
// +build js

package main

import (
	"fmt"
	"time"

	"codeberg.org/sbinet/go-web-examples/06-go-web-gopherjs/pkg"
	"github.com/gopherjs/gopherjs/js"
)

func tickTxt(t time.Duration) string {
	return fmt.Sprintf("GopherJS demo starting... (content will appear in %v)", t)
}

func main() {
	max := 5 * time.Second
	doc := js.Global.Get("document")
	doc.Get("body").Set("innerHTML", tickTxt(max))
	start := time.Now()
	timeout := time.After(max)
	ticker := time.NewTicker(10 * time.Millisecond)

loop:
	for {
		select {
		case <-timeout:
			ticker.Stop()
			break loop
		case c := <-ticker.C:
			doc.Get("body").Set("innerHTML", tickTxt(max-c.Sub(start)))
		}
	}
	str := "you"
	doc.Get("body").Set("innerHTML", pkg.Hello(str)+"\n<p>An <em>alert</em> should pop up in 3s</p>")
	<-time.After(3 * time.Second)
	js.Global.Call("alert", pkg.Hello(str))
}
