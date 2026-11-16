package server

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/geoffhotchkiss/primes.sexy/sexyprime"
)

//go:embed html
var fs embed.FS

type sexyPrimeContent struct {
	P1 int
	P2 int
}

func handler(w http.ResponseWriter, r *http.Request) {
	thing, err := template.New("index").ParseFS(fs, "html/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	tmp := sexyprime.RandomSexyPrime(sexyprime.ReadPrimes())

	content := sexyPrimeContent {
		P1: tmp,
		P2: tmp+6,
	}

	err = thing.Execute(w, content)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, fs, "html/favicon.ico")
}

func RunServer(port int) {
	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
