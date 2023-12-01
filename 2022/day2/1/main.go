package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	rounds := map[string]int{
		"A X": 4,
		"A Y": 8,
		"A Z": 3,
		"B X": 1,
		"B Y": 5,
		"B Z": 9,
		"C X": 7,
		"C Y": 2,
		"C Z": 6,
	}
	ans := 0
	for sc.Scan() {
		row := sc.Text()
		ans += rounds[row]
	}
	fmt.Println(ans)
}
