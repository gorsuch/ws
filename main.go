package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"code.google.com/p/go.net/websocket"
)

var origin string
var url string

func init() {
	flag.StringVar(&origin, "origin", "http://localhost/", "value of Origin http header")
	flag.StringVar(&url, "url", "", "websocket url we should connect to")
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()

	if url == "" {
		usage()
	}

	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		var msg = make([]byte, 512)
		var n int
		if n, err = ws.Read(msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", msg[:n])
	}
}
