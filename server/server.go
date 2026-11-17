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

type SexyPrimeHandler struct {
	SexyPrimes []int
	SexyTemplate *template.Template
}

func (h *SexyPrimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmp := sexyprime.RandomSexyPrime(h.SexyPrimes)
	content := sexyPrimeContent {
		P1: tmp,
		P2: tmp+6,
	}

	err := h.SexyTemplate.Execute(w, content)
	if err != nil {
		log.Fatal(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, fs, "html/favicon.ico")
}

func RunServer(port int) {
	thing, err := template.New("index").ParseFS(fs, "html/index.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	sexyHandler := &SexyPrimeHandler{
		SexyPrimes: sexyprime.ReadPrimes(),
		SexyTemplate: thing,
	}

	http.Handle("/", sexyHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
