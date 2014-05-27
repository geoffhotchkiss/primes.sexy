package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
)

func main() {
	primes := read_primes("primes.txt")

	for i := 0; i < 10; i++ { 
		rand_prime := random_sexy(&primes)
		fmt.Printf("(%v,%v)\n", rand_prime, rand_prime+6)
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

