package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	pass := 0
	dial := 50
	for _, line := range lines {
		dir := line[0]

		num, _ := strconv.Atoi(line[1:])
		if dir == 'L' {

			mod := num % 100
			sub := dial - mod
			if sub < 0 {
				sub += 100
			}
			dial = sub
			if dial == 0 {
				pass++
			}
		} else {
			dial += num
			dial = dial % 100
			if dial == 0 {
				pass++
			}
		}
	}
	fmt.Println(pass)

}
