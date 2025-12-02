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

		clicks := num / 100
		pass += clicks

		mod := num % 100
		if mod == 0 {
			continue
		}

		if dir == 'L' {
			distToZero := dial
			if dial == 0 {
				distToZero = 100
			}
			if mod >= distToZero {
				pass++
			}
			sub := dial - mod
			if sub < 0 {
				sub += 100
			}
			dial = sub
		} else {
			distToZero := 100 - dial

			if dial == 0 {
				distToZero = 100
			}

			if mod >= distToZero {
				pass++
			}

			dial = (dial + mod) % 100
		}
	}

	fmt.Println(pass)
}
