package main

func swapTop2(list *NBlist) *NBlist {
	var tmp *NBlist

	if list == nil || list.next == nil {
		return list
	}
	tmp = list.next
	list.next = tmp.next
	tmp.next = list
	return tmp
}

func push(dest, src **NBlist) {
	var tmp *NBlist

	if src == nil || *src == nil {
		return
	}
	if dest == nil || *dest == nil {
		*dest = extractFirst(src)
	} else {
		tmp = *dest
		*dest = extractFirst(src)
		(*dest).next = tmp
	}
}

func rotate(list *NBlist) *NBlist {
	var tmp *NBlist
	tmp = extractFirst(&list)
	return addBack(list, tmp)
}

func revRotate(list *NBlist) *NBlist {
	var tmp *NBlist
	tmp = extractLast(&list)
	return addFront(list, tmp)
}

func clearPs(ps *Pushswap) {
	clearList(&(ps.org))
	clearList(&(ps.a.lst))
	clearList(&(ps.b.lst))
	clearCmds(&(ps.cmds))
	ps.a = nil
	ps.b = nil
}

func psChecker(a, b **NBlist, cmd *Commands) {
	var ptr *Commands

	ptr = cmd
	for ptr != nil {
		cmdToFunc(a, b, ptr.abbr)
		ptr = ptr.next
	}
	isCorrectPs(*a, *b)
}

func psSortTop3A2(ps *Pushswap) int {
	if ps.a.lst.next.next.nb > ps.a.lst.next.nb && ps.a.lst.nb > ps.a.lst.next.next.nb {
		sa(ps)
		ra(ps)
		sa(ps)
		return ps.a.lst.next.nb
	} else if ps.a.lst.next.next.nb < ps.a.lst.next.nb && ps.a.lst.nb > ps.a.lst.next.nb {
		sa(ps)
		ra(ps)
		sa(ps)
		rra(ps)
		sa(ps)
		return ps.a.lst.next.next.nb
	}
	return ps.a.lst.next.next.nb
}

func psSortTop3A(ps *Pushswap) int {
	if ps.a.lst.nb < ps.a.lst.next.next.nb && ps.a.lst.next.nb > ps.a.lst.next.next.nb {
		ra(ps)
		sa(ps)
		return (ps.a.lst.next.nb)
	} else if ps.a.lst.nb > ps.a.lst.next.nb && ps.a.lst.nb < ps.a.lst.next.next.nb {
		sa(ps)
		return (ps.a.lst.next.next.nb)
	} else if ps.a.lst.nb < ps.a.lst.next.nb && ps.a.lst.nb > ps.a.lst.next.next.nb {
		ra(ps)
		sa(ps)
		rra(ps)
		sa(ps)
		return (ps.a.lst.next.next.nb)
	}
	return psSortTop3A2(ps)
}

func psASort3(ps *Pushswap) int {
	if isRotSort(ps.a) == 0 {
		sa(ps)
	}
	return findMax(ps.a.lst)
}

func psSmartRotateA(ps *Pushswap) {
	var i int
	i = findNBpos(ps.a.lst, ps.a.min)
	rotA(ps, i)
}

func psCheckSortN(list *NBlist, len int) bool {
	for list != nil && len > 1 {
		if list.next != nil && list.nb > list.next.nb {
			return false
		}
		list = list.next
		len--
	}
	return true
}

func psGetIVal(list *NBlist, i int) int {
	for j := i; j > 0; j-- {
		list = list.next
	}
	return list.nb
}

func pushSwap(ps *Pushswap) {
	var runs *NBlist
	var left int
	var right int

	psInfo(ps)
	if ps.a.count == 3 {
		psASort3(ps)
	} else if ps.a.count > 3 {
		runs = nil
		left = findMin(ps.a.lst) - 1
		right = getMedian(ps.a.lst)
		doRun(ps, left, right, &runs)
	}
	psSmartRotateA(ps)
	parseCmds(&(ps.cmds))
	printCmdList(ps.cmds)
	clearPs(ps)
	clearList(&runs)
}
