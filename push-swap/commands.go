package main

import (
	"fmt"
)

func createCmdNode(abbr *string) *Commands {
	node := &Commands{}
	node.abbr = abbr
	node.next = nil
	return node
}

func addCmdFront(head, node *Commands) *Commands {
	if head == nil {
		return node
	}
	node.next = head
	return node
}

func addCmdBack(head, node *Commands) *Commands {
	var tmp *Commands

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

func createCmdFront(head *Commands, abbr *string) *Commands {
	var node *Commands

	node = createCmdNode(abbr)

	return addCmdFront(head, node)
}

func createCmdBack(head *Commands, abbr *string) *Commands {
	var node *Commands

	node = createCmdNode(abbr)

	return addCmdBack(head, node)
}

func countCmdList(list *Commands) int {
	var cnt int

	cnt = 0

	for list != nil {
		cnt++
		list = list.next
	}

	return 0
}

func freeCmd(node *Commands) {
	node.abbr = nil
	node = nil
}

func removeCmd(cmd **Commands) {
	var node *Commands

	node = (*cmd).next
	(*cmd).next = (*cmd).next.next
	node.next = nil
	node = nil
}

func findRB(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd

	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "ra" {
		findRB(&(tmp.next))
	}
	for tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rr" {
		tmp = tmp.next
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rb" {
		(*cmd).abbr = nil
		str := "rr"
		(*cmd).abbr = &str
		removeCmd(&tmp)
	}
}

func parseRA(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "ra" {
			findRB(&tmp)
		}
		tmp = tmp.next
	}
}

func findRRB(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd

	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rra" {
		findRRB(&(tmp.next))
	}
	for tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rrr" {
		tmp = tmp.next
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rrb" {
		(*cmd).abbr = nil
		str := "rrr"
		(*cmd).abbr = &str
		removeCmd(&tmp)
	}
}

func parseRRA(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd

	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "rra" {
			findRRB(&tmp)
		}
		tmp = tmp.next
	}
}

func findRA(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rb" {
		findRA(&(tmp.next))
	}

	for tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rr" {
		tmp = tmp.next
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "ra" {
		(*cmd).abbr = nil
		str := "rr"
		(*cmd).abbr = &str
		removeCmd(&tmp)
	}
}

func parseRB(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd

	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "rb" {
			findRA(&tmp)
		}
		tmp = tmp.next
	}
}

func findRRA(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rrb" {
		findRRA(&(tmp.next))
	}

	for tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rrr" {
		tmp = tmp.next
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rra" {
		(*cmd).abbr = nil
		str := "rrr"
		(*cmd).abbr = &str
		removeCmd(&tmp)
	}
}

func parseRRB(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd

	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "rrb" {
			findRRA(&tmp)
		}
		tmp = tmp.next
	}
}

func parseCmds(cmd **Commands) {
	parseRA(cmd)
	parseRB(cmd)
	parseRRA(cmd)
	parseRRB(cmd)
}

func findRRArot(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "ra" {
		findRRArot(&(tmp.next))
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rra" {
		removeCmd(&(tmp.next))
		removeCmd(&tmp)
	}
}

func parseRArot(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "ra" {
			findRRArot(&tmp)
		}
		tmp = tmp.next
	}
}

func findRArot(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "rra" {
		findRArot(&(tmp.next))
	}
	if tmp.next != nil && tmp.next.abbr != nil && *tmp.next.abbr == "ra" {
		removeCmd(&(tmp.next))
		removeCmd(&tmp)
	}
}

func parseRRArot(cmd **Commands) {
	var tmp *Commands

	tmp = *cmd
	for tmp != nil {
		if tmp.abbr != nil && *tmp.abbr == "rra" {
			findRArot(&tmp)
		}
		tmp = tmp.next
	}
}

func parseRotates(cmd **Commands) {
	parseRArot(cmd)
	parseRRArot(cmd)
}

func cmdToFunc(a **NBlist, b **NBlist, cmd *string) {
	if *cmd == "sa" {
		*a = swapTop2(*a)
	} else if *cmd == "sb" {
		*b = swapTop2(*b)
	} else if *cmd == "ss" {
		*a = swapTop2(*a)
		*b = swapTop2(*b)
	} else if *cmd == "pa" {
		push(a, b)
	} else if *cmd == "pb" {
		push(b, a)
	} else if *cmd == "ra" {
		*a = rotate(*a)
	} else if *cmd == "rb" {
		*b = rotate(*b)
	} else {
		cmdToFunc2(a, b, cmd)
	}
}

func cmdToFunc2(a **NBlist, b **NBlist, cmd *string) {
	if *cmd == "rr" {
		*a = rotate(*a)
		*b = rotate(*b)
	} else if *cmd == "rra" {
		*a = revRotate(*a)
	} else if *cmd == "rrb" {
		*b = revRotate(*b)
	} else if *cmd == "rrr" {
		*a = revRotate(*a)
		*b = revRotate(*b)
	} else {
		errorExit("Unknown Command!")
	}
}

func extractFirstCmd(head **Commands) *Commands {
	var first *Commands

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

func clearCmds(cmds **Commands) {
	var tmp *Commands
	for tmp = extractFirstCmd(cmds); tmp != nil; tmp = extractFirstCmd(cmds) {
		tmp.abbr = nil
		tmp = nil
	}
}

func printCmdList(list *Commands) {
	for list != nil {
		if list.abbr != nil {
			fmt.Println(*list.abbr)
		}
		list = list.next
	}
	return
}
