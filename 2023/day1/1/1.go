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

	digits := "0123456789"

	ans := 0
	for sc.Scan() {
		row := sc.Text()

		t := ""

		for i := 0; i < len(row); i++ {
			if strings.Contains(digits, string(row[i])) {
				t += string(row[i])
				break
			}
		}

		for i := len(row) - 1; i >= 0; i-- {
			if strings.Contains(digits, string(row[i])) {
				t += string(row[i])
				break
			}
		}

		q, _ := strconv.Atoi(t)
		ans += q
	}

	fmt.Println(ans)

}
