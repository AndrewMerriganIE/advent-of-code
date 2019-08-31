// learning project main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var id int
	var id1, id2, res string
	id, id1, id2, res = findCommonID(lines)

	fmt.Println(string(res))

	fmt.Println(string(id1))
	fmt.Println(string(id2))

	fmt.Println(string(id1[id]))
	fmt.Println(string(id2[id]))
}

type boxid struct {
	val string
}

func findCommonID(boxIDs []string) (int, string, string, string) {
	for i := 0; i < len(boxIDs[0]); i++ {
		commonIDs := make(map[string]boxid, len(boxIDs))
		for _, curID := range boxIDs {
			curSeq := curID[:i] + curID[i+1:]
			if id, ok := commonIDs[curSeq]; ok {
				return i, curID, id.val, curSeq
			}
			commonIDs[curSeq] = boxid{curID}
		}
	}
	return 0, "", "", ""
}
