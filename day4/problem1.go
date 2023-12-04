package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("[0-9]+|\\||:")
	total := 0
	for scanner.Scan() {
		sum := 0
		line := scanner.Text()

		match := r.FindAllString(line, -1)
		con := make(map[string]bool)
		flag := false
		for i := 0; i < len(match); i++ {
			if match[i] == ":" {
				flag = true
				continue
			}
			if match[i] == "|" {
				flag = false
			}
			if flag {
				con[match[i]] = true
				continue
			}

			if con[match[i]] {
				if sum == 0 {
					sum++
					continue
				}
				sum *= 2
			}
		}
		total += sum
	}
	fmt.Println(total)
}
