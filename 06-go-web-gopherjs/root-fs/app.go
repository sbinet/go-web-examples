// +build js

package main

import (
	"fmt"
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/sbinet/go-web-examples/06-go-web-gopherjs/pkg"
)

func tickTxt(t time.Duration) string {
	return fmt.Sprintf("GopherJS demo starting... (content will appear in %v)", t)
}

func main() {
	max := 3 * time.Second
	doc := js.Global.Get("document")
	doc.Get("body").Set("innerHTML", tickTxt(max))
	start := time.Now()
	timeout := time.After(max)
	ticker := time.NewTicker(100 * time.Millisecond)
loop:
	for {
		select {
		case <-timeout:
			break loop
		case c := <-ticker.C:
			doc.Get("body").Set("innerHTML", tickTxt(max-c.Sub(start)))
		}
	}
	str := "you"
	doc.Get("body").Set("innerHTML", pkg.Hello(str)+"\n<p>An <em>alert</em> should pop up in 5s</p>")
	<-time.After(5 * time.Second)
	js.Global.Call("alert", pkg.Hello(str))
}
