package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"time"
)

var addr string
var help bool
var sleep time.Duration

func main() {
	flag.StringVar(&addr, "bind", ":8080", "address to listen")
	flag.BoolVar(&help, "help", false, "help")
	flag.DurationVar(&sleep, "sleep", 10*time.Millisecond, "sleep")
	flag.Parse()

	status := fmt.Sprintf("--- webup %s sleep(%s)", addr, sleep)

	if help {
		flag.Usage()
		os.Exit(0)
	}

	f := func(w http.ResponseWriter, r *http.Request) {
		b, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf(">>>\n%s", b)

		time.Sleep(sleep)

		rec := httptest.NewRecorder()
		h := http.FileServer(http.Dir("."))
		h.ServeHTTP(rec, r)

		_, _ = fmt.Fprint(w, rec.Body)

		b, err = httputil.DumpResponse(rec.Result(), true)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("<<<\n%s", b)
		log.Println(status)
	}

	log.Println(status)
	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(f)))
}
