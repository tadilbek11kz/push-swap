package main

import (
	"strconv"
	"strings"
)

func fetchNumber(ps *Pushswap, ac int, av []string) {
	readArgs(ps, ac, av)
}

func nbParseStr(ps *Pushswap, str string) {
	nums := strings.Split(str, " ")

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			errorExit("List contains invalid characters")
		}
		ps.a.lst = createBack(ps.a.lst, n)
		ps.org = createBack(ps.org, n)
	}
}

func readArgs(ps *Pushswap, ac int, av []string) {
	args := strings.Join(av, " ")
	nbParseStr(ps, args)
}
