package main

import (
	"fmt"
	"os"
)

func errorExit(s string) {
	fmt.Printf("error: %s\n", s)
	os.Exit(1)
}

func printList(list *NBlist, name string, sep string) {
	fmt.Printf("%v: ", name)
	if list == nil {
		fmt.Println()
		return
	}
	for list.next != nil {
		fmt.Printf("%v%v", list.nb, sep)
		list = list.next
	}
	fmt.Printf("%v\n", list.nb)
	return
}
