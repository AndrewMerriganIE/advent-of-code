package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	min_length := math.MaxInt64
	fmt.Println(min_length)

	for expectedChar := 65; expectedChar < 65+26; expectedChar++ {

		file, err := os.Open("input.txt")
		if err != nil {
			os.Exit(2)
		}

		defer file.Close()

		reader := bufio.NewReader(file)

		firstCharacter, _, _ := reader.ReadRune()
		possibleFirstCharacter := firstCharacter
		for possibleFirstCharacter != 0 {
			if possibleFirstCharacter != rune(expectedChar) && possibleFirstCharacter != rune(expectedChar+32) {
				firstCharacter = possibleFirstCharacter
				break
			}
			possibleFirstCharacter, _, _ = reader.ReadRune()
		}

		seenCharacters := []rune{possibleFirstCharacter}
		currentCharacter, _, _ := reader.ReadRune()

		for currentCharacter != 0 {
			if currentCharacter != rune(expectedChar) && currentCharacter != rune(expectedChar+32) {
				seenCharacters = processRune(currentCharacter, seenCharacters, rune(expectedChar))
			}
			currentCharacter, _, _ = reader.ReadRune()
		}

		current_length := len(seenCharacters)

		if current_length < min_length {
			min_length = current_length
		}
	}
	fmt.Println(min_length)
}

func processRune(current rune, seen []rune, expectedChar rune) []rune {

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
		return processRune(seen[len(seen)-2], seen[:len(seen)-2], expectedChar)

	} else {
		//fmt.Println("returning:: ")
		return append(seen, current)
	}
}
