package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	ans := 0
	for sc.Scan() {
		row := sc.Text()
		ranges := strings.Split(row, ",")
		x := strings.Split(ranges[0], "-")
		y := strings.Split(ranges[1], "-")
		a, _ := strconv.Atoi(x[0])
		b, _ := strconv.Atoi(x[1])
		c, _ := strconv.Atoi(y[0])
		d, _ := strconv.Atoi(y[1])
		if !(c-b > 0 && a-d < 0 || c-b < 0 && a-d > 0) {
			ans++
		}
	}
	fmt.Println(ans)
}
