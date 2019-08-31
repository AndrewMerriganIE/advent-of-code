package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

type Input struct {
	N int
	X float64
	Y float64
}

func readInput(filename string) []Input {
	fileHandle, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	inputList := []Input{}

	n := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) > 0 {
			i := Input{}

			if _, err := fmt.Sscanf(line, "%b, %b", &i.X, &i.Y); err != nil {
				log.Fatal(err)
			}
			i.N = n
			n++

			inputList = append(inputList, i)
		}
	}

	return inputList
}

var file = flag.String("file", "./input.txt", "file used for input")

func main() {
	flag.Parse()

	inputList := readInput(*file)

	fmt.Println(part1(inputList))
}

type coord struct {
	X float64
	Y float64
}

func part1(inputList []Input) (int, int) {
	height := float64(0)
	width := float64(0)
	maxSum := float64(10000)

	inf := make(map[coord]bool)
	m := make(map[coord]int)
	var coords []coord

	regions := 0

	for _, input := range inputList {
		if input.Y >= height {
			height = input.Y
		}

		if input.X >= width {
			width = input.X
		}

		coords = append(coords, coord{X: input.X, Y: input.Y})
	}

	for y := float64(0); y <= height; y++ {
		for x := float64(0); x <= width; x++ {
			mc := coord{0, 0}
			min := float64(-1)
			tot := float64(0)
			for _, c := range coords {
				dist := math.Abs(x-c.X) + math.Abs(y-c.Y)
				if dist < min || min == -1 {
					min = dist
					mc = c
				} else if dist == min {
					mc = coord{-1, -1}
				}

				// Part 2
				tot += math.Abs(x-c.X) + math.Abs(y-c.Y)
			}

			if x == 0 || y == 0 || x == width || y == height {
				inf[mc] = true
			}

			m[mc]++

			// Part 2
			if tot < maxSum {
				regions++
			}
		}
	}

	max := 0
	for k, v := range m {
		if _, found := inf[k]; v > max && !found {
			max = v
		}
	}

	return max, regions
}
