package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
	"net/http"
	"http/fcgi"
	"net"
	"html/template"
	"time"
)

type SexyPrime struct {
	P1 int64
	P2 int64
}

const (
	SexyStamp = "2006/1/2 15:04:05.000000000" 
)

func main() {
	primes := read_primes("/home/geoff/projects/go/src/github.com/geoffhotchkiss/primes.sexy/primes.txt")
	seed := time.Now().Unix()
	rand.Seed(seed)

	fmt.Println("Starting server...")
	fmt.Println("Seed:", seed)
	fmt.Println()

	http.HandleFunc("/", primeHandler(primes))
	serveSingle("/favicon.ico","./favicon.ico")
	serveSingle("/html/sexy.css","./html/sexy.css")
	http.ListenAndServe(":8081", nil)
}

func primeHandler(p []int64) func(http.ResponseWriter, *http.Request) {
	return func(wt http.ResponseWriter, rt *http.Request) {
		sexy_prime := random_sexy(&p)
		sp := SexyPrime{P1: sexy_prime, P2: sexy_prime+6}
		currentTime := time.Now().Format(SexyStamp)

		t, _ := template.ParseFiles("html/index.html")
		t.Execute(wt, sp)

		fmt.Printf("%v http request: %v\n", currentTime, rt.RemoteAddr)
	}
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func random_sexy(primes *[]int64) int64 {
	rand_int := rand.Intn(len(*primes))
	rand_prime := (*primes)[rand_int]

	return rand_prime
}

func read_primes(filename string) []int64 {
	file, err := os.Open(filename)

	if err != nil {
			panic(fmt.Sprintf("Got error: %v\n", err))
	}

	primes := make([]int64, 0)
	scanner := bufio.NewScanner(file)	

	for scanner.Scan() {
		line := scanner.Text()

		prime, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		primes = append(primes,prime)
	}

	return primes
}

