package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var a []int
	var b []int
	args := os.Args[1:]
	if len(args) < 1 {
		return
	}
	args = strings.Split(strings.Join(args, " "), " ")
	for _, num := range args {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error")
			return
		}
		a = append(a, n)
	}
	var commands []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "" {
			commands = append(commands, scanner.Text())
		}
	}

	for _, command := range commands {
		switch cmd := command; cmd {
		case "sa":
			a[0], a[1] = a[1], a[0]
		case "sb":
			b[0], b[1] = b[1], b[0]
		case "ss":
			a[0], a[1] = a[1], a[0]
			b[0], b[1] = b[1], b[0]
		case "pa":
			a = append([]int{b[0]}, a...)
			b = b[1:]
		case "pb":
			b = append([]int{a[0]}, b...)
			a = a[1:]
		case "ra":
			tmp := a[0]
			a = a[1:]
			a = append(a, tmp)
		case "rb":
			tmp := b[0]
			b = b[1:]
			b = append(b, tmp)
		case "rr":
			tmp := a[0]
			a = a[1:]
			a = append(a, tmp)
			tmp = b[0]
			b = b[1:]
			b = append(b, tmp)
		case "rra":
			tmp := a[len(a)-1]
			a = a[:len(a)-1]
			a = append([]int{tmp}, a...)
		case "rrb":
			tmp := b[len(b)-1]
			b = b[:len(b)-1]
			b = append([]int{tmp}, b...)
		case "rrr":
			tmp := a[len(a)-1]
			a = a[:len(a)-1]
			a = append([]int{tmp}, a...)
			tmp = b[len(b)-1]
			b = b[:len(b)-1]
			b = append([]int{tmp}, b...)
		default:
			fmt.Println("Error")
			return
		}
	}
	if sort.IntsAreSorted(a) && len(b) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
