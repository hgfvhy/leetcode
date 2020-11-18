package main

import (
	"fmt"
	"strings"
)

func main()  {
	convert("djadjsdhflksjdhdfkgjghkdfj",4)
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	n := 2 * numRows - 2
	for i, char := range s {
		fmt.Println(char)
		x := i % n
		rows[min(x, n - x)] += string(char)
	}
	return strings.Join(rows, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}