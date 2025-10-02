package main

import "fmt"

func ReplaceBytesNum(a, i int64) int64 {
	return a ^ (1<<i - 1)
}

func main() {
	var a, i int64
	fmt.Scan(&a, &i)
	fmt.Println(ReplaceBytesNum(a, i))
}
