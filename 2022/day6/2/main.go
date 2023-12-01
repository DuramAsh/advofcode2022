package main

import (
	"bufio"
	"fmt"
	"github.com/hashicorp/go-set"
	"os"
)

func main() {

	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
free:
	for sc.Scan() {
		row := sc.Text()
		for i, _ := range row {
			slice := make([]byte, 0)
			for j := 0; j < 14; j++ {
				slice = append(slice, row[i+j])
			}
			s := set.From[byte](slice).Slice()
			if len(s) == 14 {
				fmt.Println(i + 14)
				break free
			}
		}
	}
}
