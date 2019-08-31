// learning project main.go
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

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	file.Close()

	claims := parseClaims(lines)
	calculateOverlap := calculateOverlap(claims)
	fmt.Println(calculateOverlap)
}

type claimStruct struct {
	Id       string
	Start_x  int
	Start_y  int
	Offset_x int
	Offset_y int
}

func parseClaims(raw_claims []string) []claimStruct {

	var claims []claimStruct
	r, _ := regexp.Compile("(#[0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	for _, rawClaim := range raw_claims {
		split_claim := r.FindStringSubmatch(rawClaim)
		var claim claimStruct
		claim.Id = split_claim[1]
		num1, _ := strconv.Atoi(split_claim[2])
		claim.Start_x = num1
		num2, _ := strconv.Atoi(split_claim[3])
		claim.Start_y = num2
		num3, _ := strconv.Atoi(split_claim[4])
		claim.Offset_x = num3
		num4, _ := strconv.Atoi(split_claim[5])
		claim.Offset_y = num4
		claims = append(claims, claim)
	}
	return claims
}

func calculateOverlap(claims []claimStruct) int {
	overlap := 0
	var map_of_fabric = map[int]map[int]int{}

	for _, claim := range claims {
		for position_y := claim.Start_y; position_y < claim.Start_y+claim.Offset_y; position_y++ {
			for position_x := claim.Start_x; position_x < claim.Start_x+claim.Offset_x; position_x++ {
				if map_of_fabric[position_x] == nil {
					map_of_fabric[position_x] = map[int]int{}
				}

				map_of_fabric[position_x][position_y] = map_of_fabric[position_x][position_y] + 1
				if map_of_fabric[position_x][position_y] == 2 {
					overlap = overlap + 1
				}
			}
		}
	}

	return overlap
}
