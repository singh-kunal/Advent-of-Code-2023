package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var lines []string
	reader := bufio.NewReader(os.Stdin)

	ref := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("err: ", err)
			return
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	fmt.Println("Result:")
	var res int = 0

	for _, line := range lines {
		str := ""
		flag := false
		for i := 0; i < len(line); i++ {
			if '0' <= line[i] && line[i] <= '9' {
				str += string(line[i])
				break
			}
			for j := i + 1; j < len(line); j++ {
				substr := line[i:j]
				if val, found := ref[substr]; found {
					str += val
					flag = true
					break
				}

			}
			if flag {
				flag = false
				break
			}

		}
		for i := len(line) - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				str += string(line[i])
				break
			}
			for j := i - 1; j >= 0; j-- {
				substr := line[j : i+1]
				if val, found := ref[substr]; found {
					str += val
					flag = true
					break
				}
			}
			if flag {
				flag = false
				break
			}

		}
		num, _ := strconv.Atoi(str)
		res += num

	}
	fmt.Println(res)
}
