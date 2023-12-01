package main

import (
	"bufio"
	"fmt"
	"github.com/duramash/go-data-structures/stack"
	"os"
	"strconv"
)

func main() {
	var stacks []*stack.Stack[byte]
	for i := 0; i < 10; i++ {
		newStack := stack.MakeStack[byte]()
		stacks = append(stacks, &newStack)
	}
	var S int
	data, _ := os.Open("./data.txt")
	defer data.Close()
	sc := bufio.NewScanner(data)
	for sc.Scan() {
		row := sc.Text()
		if len(row) == 0 {
			for i := 0; i < S; i++ {
				temp := stack.MakeStack[byte]()
				for !stacks[i].IsEmpty() {
					val, _ := stacks[i].Pop()
					temp.Push(val.Val)
				}
				stacks[i] = &temp
			}
			continue
		}
		if row[0] == 'm' {
			var from, to uint8
			var amount string
			if row[6] != ' ' {
				amount = row[5:7]
				from, to = row[13], row[18]
			} else {
				amount = string(row[5])
				from, to = row[12], row[17]
			}
			n, _ := strconv.Atoi(amount)
			f, _ := strconv.Atoi(string(from))
			t, _ := strconv.Atoi(string(to))
			for i := 0; i < n; i++ {
				val, _ := stacks[f-1].Pop()
				stacks[t-1].Push(val.Val)
			}
		} else if row[1] != '1' {
			ar := make([]byte, 0)
			for i := 1; i < len(row); i += 4 {
				ar = append(ar, row[i])
			}
			for i, el := range ar {
				S = len(ar)
				if el != ' ' {
					stacks[i].Push(el)
				}
			}
		}
	}
	for i := 0; i < S; i++ {
		fmt.Print(string(stacks[i].Peek().Val))
	}
}
