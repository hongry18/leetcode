package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(maximum69Number(9669))
}

func maximum69Number(num int) int {
	res, _ := strconv.Atoi(strings.Replace(strconv.Itoa(num), "6", "9", 1))
	return res
}
