package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"math/rand"
)

func main() {
	lines := split_lines("primes.txt")
	primes := make([]int64, len(lines))

	for i, line := range lines {
		prime, err := strconv.ParseInt(line, 10, 0)

		if err != nil {
			panic(err)
		}

		primes[i] = prime
	}


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

func split_lines(filename string) []string {
	file, err := os.Open(filename)

	if err != nil {
			panic(fmt.Sprintf("Got error: %v\n", err))
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)	

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines,line)
	}

	return lines
}

