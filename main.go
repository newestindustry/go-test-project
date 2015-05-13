package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

var listen = flag.String("listen", ":4242", "Listen address")

func init() {
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		curTime := time.Now()

		_, err := w.Write([]byte(curTime.String()))
		if err != nil {
			log.Println("Something went wrong sending the date/time back to the client:", err.Error())
		}
	})

	log.Fatal(http.ListenAndServe(*listen, nil))
}
