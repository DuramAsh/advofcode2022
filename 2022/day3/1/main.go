package main

import (
	"bufio"
	"fmt"
	"os"
)

func findIntersection(a, b string) rune {
	m := make(map[rune]bool)
	for _, el := range a {
		m[el] = true
	}
	for _, el := range b {
		if _, ok := m[el]; ok {
			return el
		}
	}
	return 'a'
}

func main() {
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	ans := 0

	for sc.Scan() {
		row := sc.Text()
		left := row[:len(row)/2]
		right := row[len(row)/2:]
		char := findIntersection(left, right)
		if 'a' <= char && char <= 'z' {
			ans += int(char) - 97 + 1
		} else {
			ans += int(char) - 65 + 27
		}
	}
	fmt.Println(ans)
}
