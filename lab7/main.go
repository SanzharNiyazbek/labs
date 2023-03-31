package main

import (
	"fmt"
	"unicode/utf8"
)

func str_len(str string) int {
	return utf8.RuneCountInString(str)
}

func main() {
	n := 0
	fmt.Scan(&n)
	array := make([]string, n)
	lengths := make([]int, n)
	for i := 0; i < n; i++ {
		var temp string
		fmt.Scan(&temp)
		array[i] = temp
	}
	for i := 0; i < n; i++ {
		lengths[i] = str_len(array[i])
		fmt.Print(lengths[i], " ")
	}
}