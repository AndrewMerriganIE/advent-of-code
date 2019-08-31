package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(2)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	firstCharacter, _, _ := reader.ReadRune()
	seenCharacters := []rune{firstCharacter}
	currentCharacter, _, _ := reader.ReadRune()

	for currentCharacter != 0 {
		seenCharacters = processRune(currentCharacter, seenCharacters)

		//
		currentCharacter, _, _ = reader.ReadRune()

		//	fmt.Println("for currentCharacter:: " + string(currentCharacter))
	}
	//fmt.Println("seenCharacters:: " + string(seenCharacters))

	fmt.Println(len(seenCharacters))
}

func processRune(current rune, seen []rune) []rune {

	if len(seen) == 0 {
		return append(seen, current)
	}

	last := seen[len(seen)-1]
	diff := current - last

	//fmt.Println("Seen::" + string(seen) + " Current::" + string(current) + "Last::" + string(last))
	//fmt.Println(int(diff))

	if diff == 32 || diff == -32 {
		// fmt.Println("TRUE")
		// fmt.Println(len(seen))
		if len(seen) < 3 {
			//fmt.Println("Seen")
			//	fmt.Println(seen[:len(seen)-1])
			return seen[:len(seen)-1]
		}
		return processRune(seen[len(seen)-2], seen[:len(seen)-2])

	} else {
		//fmt.Println("returning:: ")
		return append(seen, current)
	}
}
