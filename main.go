package main

import (
	_ "embed"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"github.com/geoffhotchkiss/primes.sexy/sexyprime"
)

//go:embed primes.txt
var primes string

func main() {
	var primes = read_primes()
	var sp = sexyprime.SexyPrime(random_sexy_prime(primes))

	fmt.Printf("%s\n", sp)
}

func random_sexy_prime(primes []int) int {
	return primes[rand.IntN(len(primes))]
}

func read_primes() []int {
	var prime_numbers []int

	for prime := range strings.FieldsSeq(primes) {
		p, err := strconv.Atoi(prime)
		if err != nil {
			fmt.Printf("Problem with converting %s to int: %v\n", prime, err)
		}
		prime_numbers = append(prime_numbers, p)
	}

	return prime_numbers
}
