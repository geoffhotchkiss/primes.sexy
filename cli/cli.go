package cli

import (
	"fmt"
	sp "github.com/geoffhotchkiss/primes.sexy/sexyprime"
	"github.com/alecthomas/kong"
)

type Context struct {
	Debug bool
}

type RandCmd struct { 
	Number int `default:"1" help:"Number of random Sexy Prime's to print."`
}

func (r *RandCmd) Run(ctx *Context) error {
	var primes = sp.ReadPrimes()

	for i := 0; i < r.Number; i++ {
		var sp = sp.SexyPrime(sp.RandomSexyPrime(primes))
		fmt.Printf("%s\n",sp)
	}

	return nil
}

type ServerCmd struct {
	Port int `default:"8088" help:"The port to listen on."`
}

func (s *ServerCmd) Run(ctx *Context) error {
	fmt.Printf("Running server on port %d\n", s.Port)
	return nil
}

var cli struct {
	Debug bool `help:"Enable debug mode."`

	Rand RandCmd `cmd:"" help:"Prints a random Sexy Prime."`
	Server ServerCmd `cmd:"" help:"Runs the primes.sexy sever."`
}

func RunCli() {
	ctx := kong.Parse(&cli)
	err := ctx.Run(&Context{Debug: cli.Debug})
	ctx.FatalIfErrorf(err)
}
