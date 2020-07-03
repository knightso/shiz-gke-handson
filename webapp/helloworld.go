package main

import (
    "flag"
    "fmt"
	"io"
	"log"
	"net/http"
)
var port int

func init() {
	flag.IntVar(&port, "p", 8080, "port number")
}

func main() {
    flag.Parse()
    
    fmt.Printf("listening port:%d\n", port)

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/", helloHandler)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal(err)
	}
}

