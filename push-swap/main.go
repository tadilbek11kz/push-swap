package main

import (
	"os"
)

func main() {
	var ps Pushswap
	args := os.Args[1:]
	if len(args) < 1 {
		errorExit("Not enough args")
	}
	ps.a = &Stack{}
	//ps.A.Lst = &NBlist{}
	ps.b = &Stack{}
	//ps.B.Lst = &NBlist{}
	ps.cmds = &Commands{}
	fetchNumber(&ps, len(args), args)
	if checkValid(ps.a.lst) == 0 {
		errorExit("Ivalid list of numbers (contains duplicates)")
	}
	pushSwap(&ps)
}
