package main

import (
	"fmt"
	"github.com/geoffhotchkiss/primes.sexy/sexyprime"
)


func main() {
	var primes = sexyprime.ReadPrimes()
	var sp = sexyprime.SexyPrime(sexyprime.RandomSexyPrime(primes))

	fmt.Printf("%s\n", sp)
}


