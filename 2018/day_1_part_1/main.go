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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err == nil {
			frequency += num
		}

	}
	fmt.Println(frequency)
}
