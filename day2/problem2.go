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
	ref := map[string]string{"red": "red", "blue": "blue", "green": "green"}

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	r, _ := regexp.Compile("red|blue|green|[0-9]+")
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		match := r.FindAllString(line, -1)

		red, green, blue := 0, 0, 0

		for i := 0; i < len(match); i++ {
			if ref[match[i]] == "red" {
				tmp, _ := strconv.Atoi(match[i-1])
				if tmp > red {
					red = tmp
				}
			}
			if ref[match[i]] == "green" {
				tmp, _ := strconv.Atoi(match[i-1])

				if tmp > green {
					green = tmp
				}
			}
			if ref[match[i]] == "blue" {
				tmp, _ := strconv.Atoi(match[i-1])

				if tmp > blue {
					blue = tmp
				}
			}

		}
		sum += red * blue * green

	}
	fmt.Println(sum)

}
