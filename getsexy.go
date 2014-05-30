package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
	"net/http"
	"html/template"
)

type SexyPrime struct {
	P1 int64
	P2 int64
}

func main() {
	primes := read_primes("/home/geoff/projects/go/src/github.com/geoffhotchkiss/primes.sexy/primes.txt")

	for i := 0; i < 10; i++ { 
		rand_prime := random_sexy(&primes)
		fmt.Printf("(%v,%v)\n", rand_prime, rand_prime+6)
	}

	http.HandleFunc("/", primeHandler(primes))
	http.ListenAndServe(":8080", nil)
}

func primeHandler(p []int64) func(http.ResponseWriter, *http.Request) {
	return func(wt http.ResponseWriter, rt *http.Request) {
		sexy_prime := random_sexy(&p)
		sp := SexyPrime{P1: sexy_prime, P2: sexy_prime+6}
		t, _ := template.ParseFiles("html/index.html")
		t.Execute(wt, sp)
		fmt.Printf("\nhttp request: %v\n", rt.RemoteAddr)
		//fmt.Fprintf(wt, "Hi there, I love (%v,%v)", sexy_prime, sexy_prime+6)
	}
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

