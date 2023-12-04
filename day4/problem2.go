package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	total := 0

	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("[0-9]+|\\||:")
	counter := make(map[int]int)
	var lines []string
	i := 1
	for scanner.Scan() {

		line := scanner.Text()
		lines = append(lines, line)
		counter[i] = 1
		i++
	}

	for i := 0; i < len(lines); i++ {
		match := r.FindAllString(lines[i], -1)
		con := make(map[string]bool)
		count := 0
		flag := false
		for j := 0; j < len(match); j++ {
			if match[j] == ":" {
				flag = true
				continue
			}
			if match[j] == "|" {
				flag = false
			}
			if flag {
				con[match[j]] = true
				continue
			}

			if con[match[j]] {
				count++
			}
		}
		for j := i + 2; j < count+i+2; j++ {
			counter[j] = counter[j] + (1 * counter[i+1])
		}
	}
	for i := 1; i <= len(counter); i++ {
		total += counter[i]
	}
	fmt.Println(total)
}
