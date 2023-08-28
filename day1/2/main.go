package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	elves := make([]int,  1)
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
	sort.Slice(elves, func(i, j int) bool {
		return elves[i] < elves[j]
	})
	ans := 0
	for _, el := range elves[len(elves)-3:] {
		ans += el
	}
	fmt.Println(ans)
}
