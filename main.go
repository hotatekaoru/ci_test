package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := sum("1", "2")
	fmt.Println(i)
}

func sum(s1, s2 string) int {
	i := convStrtoInt(s1)
	j := convStrtoInt(s2)
	return i + j
}

func convStrtoInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}