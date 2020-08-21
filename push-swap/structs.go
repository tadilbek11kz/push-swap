package main

// NBlist TO-DO
type NBlist struct {
	nb   int
	next *NBlist
}

// Commands TO-DO
type Commands struct {
	abbr *string
	next *Commands
}

// Stack TO-DO
type Stack struct {
	count int
	min   int
	max   int
	lst   *NBlist
}

// Flags TO-DO
type Flags struct {
	f     string
	fd    int
	v     string
	s     string
	sec   int
	t     string
	c     string
	color int
}

// Pushswap TO-DO
type Pushswap struct {
	count int
	min   int
	max   int
	org   *NBlist
	a     *Stack
	b     *Stack
	cmds  *Commands
	flags Flags
}

// Point TO-DO
type Point struct {
	x int
	y int
}

func psInfo(ps *Pushswap) {
	ps.count = countList(ps.a.lst)
	ps.min = findMin(ps.a.lst)
	ps.max = findMax(ps.a.lst)
	ps.a.count = countList(ps.a.lst)
	ps.a.min = findMin(ps.a.lst)
	ps.a.max = findMax(ps.a.lst)
	ps.b.count = countList(ps.b.lst)
	ps.b.min = findMin(ps.b.lst)
	ps.b.max = findMax(ps.b.lst)
}
