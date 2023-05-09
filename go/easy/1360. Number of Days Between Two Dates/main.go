package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	fmt.Println(daysBetweenDates("2019-06-29", "2019-06-30"))
	fmt.Println(daysBetweenDates("2020-01-15", "2019-12-31"))
}

func daysBetweenDates(date1 string, date2 string) int {
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, date1)
	t2, _ := time.Parse(layout, date2)

	dif := t1.Sub(t2)
	return int(math.Abs(dif.Hours() / 24))
}
