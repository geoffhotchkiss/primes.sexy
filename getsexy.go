package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
	"net/http"
	"html/template"
	"time"
	"encoding/json"
	"github.com/crawford/nap"
)

type SexyPrime struct {
	error string	
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

	fmt.Println("==================")
	fmt.Println("Using go 1.4")
	fmt.Println("Starting server...")
	fmt.Println("Seed:", seed)
	fmt.Println("==================")
	fmt.Println()

	http.Handle("/getsexy/", nap.HandlerFunc(getSexy(primes)))
	http.HandleFunc("/", primeHandler(primes))
	//http.HandleFunc("/getsexy/", primeReqHandler(primes))
	http.ListenAndServe(":8081", nil)
}

func pickPrimes(p []int64) func(*http.Request) (interface{}, nap.Status) {
	return func(r *http.Request) (interface{}, nap.Status){
		sexy_prime := random_sexy(&p)
		return SexyPrime{P1: sexy_prime, P2: sexy_prime+6}, nap.OK{} 
	}
}


func primeHandler(p []int64) func(http.ResponseWriter, *http.Request) {
	return func(wt http.ResponseWriter, rt *http.Request) {
		sexy_prime := random_sexy(&p)
		sp := SexyPrime{error: "", P1: sexy_prime, P2: sexy_prime+6}
		currentTime := time.Now().Format(SexyStamp)

		t, _ := template.ParseFiles("/home/geoff/projects/go/src/github.com/geoffhotchkiss/primes.sexy/html/index.html")
		t.Execute(wt, sp)

		fmt.Printf("%v http request: %+v\n", currentTime, rt.Header["X-Real-Ip"])
	}
}

func getSexy(p []int64) func(*http.Request) (interface{}, nap.Status) {
	return func(r *http.Request) (interface{}, nap.Status) {
		param := r.URL.Path[9:]
		num64, err := strconv.ParseInt(param, 10, 0)
		num := int(num64)

		if err != nil {
			return nil,nap.NotFound{fmt.Sprintf("could not convert param: %s to number", param)}
		}

		if num >= len(p) || num < 0 {
			return nil,nap.NotFound{fmt.Sprintf("request out of range of 0 and %d", len(p)-1)}
		} 

		prime := p[num]
		return SexyPrime{P1: prime, P2: prime+6}, nap.OK{}
	}
}

func primeReqHandler(primes []int64) func (http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		param := r.URL.Path[9:]
		badSexy := SexyPrime{error: "", P1: -1, P2: -1}
		//w.Header().Set("Content-Type","applicaiton/json")

		num64, err := strconv.ParseInt(param, 10, 0)
		num := int(num64)

		if err != nil || num > len(primes) {
			badSexy.error = fmt.Sprintf("request larter than %d", len(primes))
			fmt.Println(badSexy)
			sexyJson, _ := json.Marshal(badSexy)
			w.Write(sexyJson)
			return
		} 

		prime := primes[num]
		sexyPrime := SexyPrime{error: "", P1: prime, P2: prime+6}
		sexyJson, err := json.Marshal(sexyPrime)

		if err != nil {
			badSexy.error = err.Error()
			fmt.Println(badSexy)
			sexyJson, _ := json.Marshal(badSexy)
			w.Write(sexyJson)
			return
		}

		w.Write(sexyJson)
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

