package main

import (
	"fmt"
)

func isSortList(list *NBlist) bool {
	for list.next != nil {
		if list.nb > list.next.nb {
			return false
		}
		list = list.next
	}
	return true
}

func isCorrectPs(a, b *NBlist) {
	if isSortList(a) && b == nil {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}

func isRotSort(stack *Stack) int {
	var tmp *NBlist
	var i int
	var res int

	i = 1
	res = i
	tmp = stack.lst
	if tmp != nil && tmp.nb != stack.min && tmp.nb < getLast(tmp) {
		return 0
	}
	for tmp != nil && tmp.next != nil {
		if tmp.nb > tmp.next.nb {
			if tmp.nb == stack.max && tmp.next.nb == stack.min {
				res = i + 1
				tmp = tmp.next
				continue
			}
			return 0
		}
		tmp = tmp.next
		i++
	}
	return res
}

func isRevRotSort(stack *Stack) int {
	var tmp *NBlist
	var i int
	var res int

	i = 1
	res = i
	tmp = stack.lst
	if tmp != nil && tmp.nb != stack.max && tmp.nb > getLast(tmp) {
		return 0
	}
	for tmp != nil && tmp.next != nil {
		if tmp.nb < tmp.next.nb {
			if tmp.nb == stack.min && tmp.next.nb == stack.max {
				res = i + 1
				tmp = tmp.next
				continue
			}
			return 0
		}
		tmp = tmp.next
		i++
	}
	return res
}

func findSlot(list *NBlist, nb int) int {
	var pos int

	pos = 1
	for list != nil && nb > list.nb {
		pos++
		list = list.next
	}
	return pos
}
