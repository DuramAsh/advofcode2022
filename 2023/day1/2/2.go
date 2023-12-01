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

	mapa := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	ans := 0
	x := 1
	for sc.Scan() {
		row := sc.Text()

		t := ""
	outer:
		for i := 0; i < len(row); i++ {
			if strings.Contains(digits, string(row[i])) {
				t += string(row[i])
				break
			}

			for k, v := range mapa {
				if i+len(k) <= len(row) {
					if strings.Contains(row[i:i+len(k)], k) {
						t += v
						break outer
					}
				}
			}
		}

	outer2:
		for i := len(row) - 1; i >= 0; i-- {
			if strings.Contains(digits, string(row[i])) {
				t += string(row[i])
				break
			}
			for k, v := range mapa {
				if i-len(k) >= 0 {
					if strings.Contains(row[i-len(k):i+1], k) {
						t += v
						break outer2
					}
				}
			}
		}

		q, _ := strconv.Atoi(t)
		ans += q
		fmt.Println(x)
		x++
	}

	fmt.Println(ans)

}
