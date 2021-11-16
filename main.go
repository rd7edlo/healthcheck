package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = flag.String("port", "8080", "port to listen on")
var hostName string

func main() {
	flag.Parse()
	hostName, _ = os.Hostname()

	http.HandleFunc("/healthcheck", hc)

	e := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil)
	if e != nil {
		log.Fatal(e)
	}
}

func hc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	m := fmt.Sprintf("[%s]:ok", hostName)
	w.Write([]byte(m))
}
