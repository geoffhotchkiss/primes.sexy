package sexyprime

import "fmt"

type SexyPrime int

func (p SexyPrime) String() string {
	p1, p2 := p.Pair()
	return fmt.Sprintf("(%d, %d)", p1, p2)
}

func (p SexyPrime) Pair() (int, int) {
	var prime = int(p)
	return prime, prime + 6
}


