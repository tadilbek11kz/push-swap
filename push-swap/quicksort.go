package main

func bestCheck(a, b, best int) bool {
	max := a
	if b > a {
		max = b
	}
	return max < best
}

func curBest(ps *Pushswap, posA, posB, best int) int {
	if bestCheck(posA, posB, best) {
		if posA > posB {
			best = posA
		} else {
			best = posB
		}
	}
	if bestCheck(ps.a.count-posA, ps.b.count-posB, best) {
		if ps.a.count-posA > ps.b.count-posB {
			best = ps.a.count - posA
		} else {
			best = ps.b.count - posB
		}
	}
	if (posA + ps.b.count - posB) < best {
		best = posA + ps.b.count - posB
	}
	if (posB + ps.a.count - posA) < best {
		best = posB + ps.a.count - posA
	}
	return best
}

func insertOptimal(ps *Pushswap) int {
	var best int
	var bestDist int
	var posA int
	var posB int
	var tmp *NBlist

	bestDist = maxInt
	best = 0
	tmp = ps.b.lst
	for tmp != nil {
		posA = qsFindSlot(ps.a.lst, tmp.nb)
		posB = findNBpos(ps.b.lst, tmp.nb)
		if curBest(ps, posA, posB, bestDist) < bestDist {
			best = tmp.nb
			bestDist = curBest(ps, posA, posB, bestDist)
		}
		tmp = tmp.next
	}
	return best
}

func mergeRotateB(ps *Pushswap, pos int) {
	var i int

	if ps.b.count == 1 {
		return
	}
	if pos <= ps.b.count/2 {
		i = 0
		for i < pos {
			rb(ps)
			i++
		}
	} else {
		i = ps.b.count
		for i > pos {
			rrb(ps)
			i--
		}
	}
}

func bestRot(ps *Pushswap, posA, posB int) int {
	var best Point

	best.x = maxInt
	best.y = 0
	if bestCheck(posA, posB, best.x) {
		if posA > posB {
			best.x = posA
		} else {
			best.x = posB
		}
		best.y = 1
	}
	if bestCheck(ps.a.count-posA, ps.b.count-posB, best.x) {
		if ps.a.count-posA > ps.b.count-posB {
			best.x = ps.a.count - posA
		} else {
			best.x = ps.b.count - posB

		}
		best.y = 2
	}
	if (posA + ps.b.count - posB) < best.x {
		best.x = posA + ps.b.count - posB
		best.y = 3
	}
	if (posB + ps.a.count - posA) < best.x {
		best.x = posB + ps.a.count - posA
		best.y = 4
	}
	return best.y
}

func syncstacks(ps *Pushswap, a, b int) {
	var mode int

	mode = bestRot(ps, a, b)
	if mode == 0 {
		errorExit("Stacks cannot be synced @qs_syncstacks")
	} else if mode == 1 {
		raI(ps, a)
		rbI(ps, b)
	} else if mode == 2 {
		rraI(ps, a)
		rrbI(ps, b)
	} else if mode == 3 {
		raI(ps, a)
		rrbI(ps, b)
	} else if mode == 4 {
		rraI(ps, a)
		rbI(ps, b)
	}
}

func insertSort(ps *Pushswap) {
	var best int
	var a int
	var b int

	if ps.a.count != 3 {
		best = insertOptimal(ps)
	} else {
		best = ps.b.max
	}
	b = findNBpos(ps.b.lst, best)
	a = qsFindSlot(ps.a.lst, best)
	syncstacks(ps, a, b)
	pa(ps)

}

func qsFindSlot(list *NBlist, nb int) int {
	var i int

	i = 0
	if nb < findMin(list) || nb > findMax(list) {
		return findNBpos(list, findMin(list))
	}
	if nb < list.nb && nb > getLast(list) {
		return i
	}
	for list.next != nil {
		i++
		if list.nb < nb && list.next.nb > nb {
			break
		}
		list = list.next
	}
	return i
}

func splitRange(ps *Pushswap, left, right, len int) int {
	psInfo(ps)
	for contains(ps.a.lst, left, right) {
		if countList(ps.a.lst) == 3 {
			psASort3(ps)
			break
		}
		if left < ps.a.lst.nb && ps.a.lst.nb <= right {
			pb(ps)
		} else {
			ra(ps)
		}
	}
	return len - countList(ps.b.lst)
}

func merge(ps *Pushswap) int {
	var left int

	psInfo(ps)
	left = ps.b.max
	if ps.b.lst.nb < ps.a.min && ps.b.max < ps.a.min && isRotSort(ps.a) == 0 {
		pa(ps)
	}
	for ps.b.lst != nil {
		insertSort(ps)
		psInfo(ps)
	}
	return (left)
}

func runLenght(list *NBlist, left, right int) int {
	var tmp *NBlist
	var count int

	tmp = list
	count = 0
	for tmp != nil {
		if tmp.nb > left && tmp.nb <= right {
			count++
		}
		tmp = tmp.next
	}
	return count
}

func nextRun(runs **NBlist) int {
	var tmp *NBlist
	var res int

	if runs == nil || *runs == nil {
		return -1
	}
	tmp = extractFirst(runs)
	res = tmp.nb
	clearList(&tmp)
	return (res)
}

func contains(list *NBlist, left, right int) bool {
	var tmp *NBlist

	tmp = list
	for tmp != nil {
		if tmp.nb > left && tmp.nb <= right {
			return true
		}
		tmp = tmp.next
	}
	return false
}

func rotA(ps *Pushswap, i int) {
	if i <= ps.a.count/2 {
		raI(ps, i)
	} else {
		rraI(ps, i)
	}
}

func getMedianRange(list *NBlist, n int) int {
	var tmp *NBlist
	var res int

	if n == 0 {
		return 0
	}
	tmp = nil
	for list != nil && n > 0 {
		tmp = createBack(tmp, list.nb)
		list = list.next
		n--
	}
	res = getMedian(tmp)
	clearList(&tmp)
	return (res)
}

func specRuns(ps *Pushswap, toRun, left int) int {
	if toRun == 1 {
		left = ps.a.lst.nb
	} else if toRun == 2 {
		if ps.a.lst.nb > ps.a.lst.next.nb {
			sa(ps)
		}
		left = ps.a.lst.next.nb
	} else if toRun == 3 {
		left = psSortTop3A(ps)
	}
	return left
}

func doRun(ps *Pushswap, left, right int, run **NBlist) {
	var toRun int
	var pos int

	toRun = ps.a.count
	for isRotSort(ps.a) == 0 {
		if toRun < 4 {
			left = specRuns(ps, toRun, left)
		} else if psCheckSortN(ps.a.lst, toRun) {
			left = psGetIVal(ps.a.lst, toRun-1)
		} else {
			*run = createFront(*run, splitRange(ps, left, right, toRun))
			left = merge(ps)
		}
		if isRotSort(ps.a) != 0 || *run == nil {
			break
		}
		psInfo(ps)
		pos = findNBpos(ps.a.lst, left)
		if pos == ps.a.count-1 {
			pos = 0
		} else {
			pos = pos + 1
		}
		rotA(ps, pos)
		toRun = nextRun(run)
		right = getMedianRange(ps.a.lst, toRun)
	}
}
