package server

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/geoffhotchkiss/primes.sexy/sexyprime"
)

//go:embed html/index.tmpl
var indexHTML string

type sexyPrimeContent struct {
	P1 int
	P2 int
}

func handler(w http.ResponseWriter, r *http.Request) {
	thing, err := template.New("index").Parse(indexHTML)
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

func RunServer(port int) {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
