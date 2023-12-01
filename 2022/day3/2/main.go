package main

import (
	"bufio"
	"fmt"
	"os"
)

func findIntersection(arr []string) rune {
	m := make(map[rune]int)
	for _, el := range arr[0] {
		m[el] = 1
	}
	for _, el := range arr[1] {
		if _, ok := m[el]; ok {
			m[el] = 2
		}
	}
	for _, el := range arr[2] {
		if v, ok := m[el]; ok && v == 2 {
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
	i := 0
	var group []string
	for sc.Scan() {
		row := sc.Text()
		i++
		group = append(group, row)
		if i%3 != 0 {
			continue
		}
		char := findIntersection(group)
		if 'a' <= char && char <= 'z' {
			ans += int(char) - 97 + 1
		} else {
			ans += int(char) - 65 + 27
		}
		group = []string{}
	}

	fmt.Println(ans)
}
