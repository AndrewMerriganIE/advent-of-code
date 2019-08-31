// learning project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	number_of_twos := 0
	number_of_threes := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seen_once := make(map[rune]bool)
		seen_twice := make(map[rune]bool)
		seen_three := make(map[rune]bool)
		seen_more := make(map[rune]bool)

		for _, code := range scanner.Text() {
			if seen_more[code] {
				continue

			} else if seen_three[code] {
				seen_more[code] = true
				delete(seen_three, code)

			} else if seen_twice[code] {
				seen_three[code] = true
				delete(seen_twice, code)

			} else if seen_once[code] {
				seen_twice[code] = true

			} else {
				seen_once[code] = true
			}
		}

		if len(seen_three) > 0 {
			number_of_threes++
		}

		if len(seen_twice) > 0 {
			number_of_twos++
		}
	}

	fmt.Println(number_of_twos * number_of_threes)

	file.Close()
}
