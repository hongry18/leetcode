package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
	fmt.Println(dayOfYear("2019-02-10"))
	fmt.Println(dayOfYear("2018-03-01"))
	fmt.Println(dayOfYear("2016-03-01"))
	fmt.Println(dayOfYear("2012-03-01"))
	fmt.Println(dayOfYear("2008-03-01"))
	fmt.Println(dayOfYear("2000-03-01"))
	fmt.Println(dayOfYear("1900-03-01"))
}

func dayOfYear(date string) int {
	split := strings.Split(date, "-")
	months := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, _ := strconv.Atoi(split[0])
	m, _ := strconv.Atoi(split[1])
	d, _ := strconv.Atoi(split[2])

	if (y%4 == 0 && y%100 != 0) || y%400 == 0 {
		months[2] = 29
	}

	for i := 0; i < m; i++ {
		d += months[i]
	}

	return d
}
