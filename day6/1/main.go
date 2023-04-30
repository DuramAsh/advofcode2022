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
			slice := []byte{row[i], row[i+1], row[i+2], row[i+3]}
			s := set.From[byte](slice).Slice()
			if len(s) == 4 {
				fmt.Println(i + 4)
				break free
			}
		}
	}
}
