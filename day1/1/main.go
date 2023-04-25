package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	elves := make([]int, 1)
	i := 0
	for sc.Scan() {
		text := sc.Text()
		if len(text) == 0 {
			i++
			elves = append(elves, 0)
		} else {
			calories, _ := strconv.Atoi(text)
			elves[i] += calories
		}
	}
	max := math.MinInt
	for _, el := range elves {
		if max < el {
			max = el
		}
	}
	fmt.Println(max)
}
