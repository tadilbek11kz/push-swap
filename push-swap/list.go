package main

const maxUint = ^uint(0)
const minUint = 0
const maxInt = int(maxUint >> 1)
const minInt = -maxInt - 1

func createNode(nb int) *NBlist {
	node := &NBlist{}
	node.nb = nb
	node.next = nil
	return node
}

func addFront(head *NBlist, node *NBlist) *NBlist {
	if head == nil {
		return node
	}
	node.next = head
	return node
}

func addBack(head *NBlist, node *NBlist) *NBlist {
	var tmp *NBlist

	if head == nil {
		return node
	}
	tmp = head
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = node
	return head
}

func createFront(head *NBlist, nb int) *NBlist {
	var node *NBlist
	node = createNode(nb)
	return addFront(head, node)
}

func createBack(head *NBlist, nb int) *NBlist {
	var node *NBlist
	node = createNode(nb)
	return addBack(head, node)
}

/*
** Extract First
 */

func extractFirst(head **NBlist) *NBlist {
	var first *NBlist

	if head == nil || *head == nil {
		return nil
	}
	if (*head).next == nil {
		first = *head
		*head = nil
	} else {
		first = *head
		*head = (*head).next
		first.next = nil
	}
	return first
}

/*
** Extract Last
 */

func extractLast(head **NBlist) *NBlist {
	var current *NBlist
	var last *NBlist

	if head == nil || *head == nil {
		return nil
	}
	if (*head).next == nil {
		last = *head
		*head = nil
	} else {
		current = *head
		for current.next.next != nil {
			current = current.next
		}
		last = current.next
		current.next = nil
	}
	return last
}

func findMin(list *NBlist) int {
	var min int

	min = maxInt
	for list != nil {
		if min > list.nb {
			min = list.nb
		}
		list = list.next
	}
	return min
}

func findMax(list *NBlist) int {
	var max int

	max = minInt
	for list != nil {
		if max < list.nb {
			max = list.nb
		}
		list = list.next
	}
	return max
}

func countList(list *NBlist) int {
	cnt := 0
	for list != nil {
		cnt++
		list = list.next
	}
	return cnt
}

func getLast(list *NBlist) int {
	for list.next != nil {
		list = list.next
	}
	return list.nb
}

func checkValid(list *NBlist) int {
	var tmp *NBlist

	for list != nil {
		tmp = list.next
		for tmp != nil {
			if list.nb == tmp.nb {
				return 0
			}
			tmp = tmp.next
		}
		list = list.next
	}

	return 1
}

func findNBpos(list *NBlist, nb int) int {
	i := 0
	for list != nil && list.nb != nb {
		list = list.next
		i++
	}
	if list != nil {
		return i
	}
	return -1
}

func findSlotRotsort(list *NBlist, nb, max, min int) int {
	var i int

	if nb > max {
		return findNBpos(list, max) + 1
	} else if nb < min {
		return findNBpos(list, min)
	}
	i = 0
	for list != nil && list.next != nil {
		if nb > list.nb && nb < list.next.nb {
			return i
		}
		list = list.next
		i++
	}
	return 0
}

func findSlotRevRotsort(list *NBlist, nb, max, min int) int {
	var i int

	if nb > max {
		return findNBpos(list, max)
	} else if nb < min {
		return findNBpos(list, min) + 1
	}

	i = 0

	if list == nil || (nb > list.nb && nb < getLast(list)) {
		return i
	}
	for list != nil && list.next != nil {
		i++
		if nb < list.nb && nb > list.next.nb {
			return i
		}
		list = list.next
	}

	return i
}

func copyList(list *NBlist) *NBlist {
	var res *NBlist

	res = nil
	for list != nil {
		res = createBack(res, list.nb)
		list = list.next
	}

	return res
}

func extractNB(head **NBlist, nb int) *NBlist {
	var node *NBlist
	var current *NBlist

	if *head == nil {
		return nil
	}
	current = *head
	if current.nb == nb {
		return extractFirst(head)
	}
	for current.next != nil {
		if current.next.nb == nb {
			node = current.next
			current.next = node.next
			node.next = nil
			return node
		}
		current = current.next
	}
	return nil
}

func psJoinLsts(a, b *NBlist) *NBlist {
	var tmp *NBlist

	tmp = a
	if tmp == nil {
		return b
	}
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = b
	return a
}

// Not sure
func clearList(list **NBlist) {
	var tmp *NBlist

	for tmp = extractFirst(list); tmp != nil; tmp = extractFirst(list) {
		tmp = nil
	}
}

// I think i can delete clearlist func
func getMedian(list *NBlist) int {
	var tmp *NBlist
	var nb int
	var node *NBlist

	tmp = copyList(list)
	for countList(tmp) > 1 {
		nb = findMin(tmp)
		node = extractNB(&tmp, nb)
		clearList(&node)
		if countList(tmp) == 1 {
			break
		}
		nb = findMax(tmp)
		node = extractNB(&tmp, nb)
		clearList(&node)
	}
	nb = tmp.nb
	clearList(&tmp)
	return nb
}
