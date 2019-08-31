// learning project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	frequency := 0
	seen := make(map[int]bool)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

Found:
	for {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			num, _ := strconv.Atoi(scanner.Text())
			frequency += num
			if seen[frequency] {
				fmt.Printf("Result: %d \n", frequency)
				break Found
			}
			seen[frequency] = true
		}
	}
	file.Close()
}
