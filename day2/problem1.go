package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("red|blue|green|[0-9]+")
	gameno := 0
	for scanner.Scan() {
		line := scanner.Text()

		match := r.FindAllString(line, -1)
		temp, _ := strconv.Atoi(match[0])
		gameno += temp

		for i := 0; i < len(match); i++ {
			if match[i] == "red" {
				tmp, _ := strconv.Atoi(match[i-1])
				if tmp > 12 {
					gameno -= temp
					break
				}
			}
			if match[i] == "green" {
				tmp, _ := strconv.Atoi(match[i-1])
				if tmp > 13 {
					gameno -= temp
					break
				}
			}
			if match[i] == "blue" {
				tmp, _ := strconv.Atoi(match[i-1])
				if tmp > 14 {
					gameno -= temp
					break
				}
			}
		}
	}
	fmt.Println(gameno)

}
