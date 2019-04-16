package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	var addr string
	var help bool
	var sleep int
	flag.StringVar(&addr, "bind", ":8080", "address to listen")
	flag.BoolVar(&help, "help", false, "help")
	flag.IntVar(&sleep, "sleep", 0, "ms")
	flag.Parse()

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
		log.Printf("%s", b)

		h := http.FileServer(http.Dir("."))
		h.ServeHTTP(w, r)
	}

	log.Fatal(http.ListenAndServe(addr, http.HandlerFunc(f)))
}
