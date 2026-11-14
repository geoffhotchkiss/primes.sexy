package sexyprime

import (
	"fmt"
	_ "embed"
	"math/rand/v2"
	"strings"
	"strconv"
)

//go:embed primes.txt
var primes string

type SexyPrime int

func (p SexyPrime) String() string {
	p1, p2 := p.Pair()
	return fmt.Sprintf("(%d, %d)", p1, p2)
}

func (p SexyPrime) Pair() (int, int) {
	var prime = int(p)
	return prime, prime + 6
}

func RandomSexyPrime(primes []int) int {
	return primes[rand.IntN(len(primes))]
}

func ReadPrimes() []int {
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
