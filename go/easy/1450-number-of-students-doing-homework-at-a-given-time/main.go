package main

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var cnt int

	for i, t := range startTime {
		if t > queryTime || endTime[i] < queryTime {
			continue
		}
		cnt++
	}
	return cnt
}
