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
		for _, ch := range line {
			if '0' <= ch && ch <= '9' {
				str += string(ch)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if '0' <= line[i] && line[i] <= '9' {
				str += string(line[i])
				break
			}
		}
		num, _ := strconv.Atoi(str)
		res += num

	}
	fmt.Println(res)
}
