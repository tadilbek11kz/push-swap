package main

func pb(ps *Pushswap) {
	push(&(ps.b.lst), &(ps.a.lst))
	cmd := "pb"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
	psInfo(ps)
}

func pa(ps *Pushswap) {
	push(&(ps.a.lst), &(ps.b.lst))
	cmd := "pa"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
	psInfo(ps)
}

func sa(ps *Pushswap) {
	ps.a.lst = swapTop2(ps.a.lst)
	cmd := "sa"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func sb(ps *Pushswap) {
	ps.b.lst = swapTop2(ps.b.lst)
	cmd := "sb"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func ra(ps *Pushswap) {
	ps.a.lst = rotate(ps.a.lst)
	cmd := "ra"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func rb(ps *Pushswap) {
	ps.b.lst = rotate(ps.b.lst)
	cmd := "rb"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func rra(ps *Pushswap) {
	ps.a.lst = revRotate(ps.a.lst)
	cmd := "rra"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func rrb(ps *Pushswap) {
	ps.b.lst = revRotate(ps.b.lst)
	cmd := "rrb"
	ps.cmds = createCmdBack(ps.cmds, &cmd)
}

func raI(ps *Pushswap, i int) {
	for cnt := 0; cnt < i; cnt++ {
		ra(ps)
	}
}

func rbI(ps *Pushswap, i int) {
	for cnt := 0; cnt < i; cnt++ {
		rb(ps)
	}
}

func rraI(ps *Pushswap, i int) {
	for cnt := i; cnt < ps.a.count; cnt++ {
		rra(ps)
	}
}

func rrbI(ps *Pushswap, i int) {
	for cnt := i; cnt < ps.b.count; cnt++ {
		rrb(ps)
	}
}
