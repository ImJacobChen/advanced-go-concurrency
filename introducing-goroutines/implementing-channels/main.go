package main

import (
	"fmt"
	"runtime"
	"strings"
)

var loremIpsum string
var finalIpsum string
var letterSentChan chan string

func deliverToFinal(letter string, finalIpsum *string) {
	*finalIpsum += letter
}

func capitalize(current *int, length int, letters []byte, finalIpsum *string) {
	for *current < length {
		thisLetter := strings.ToUpper(string(letters[*current]))
		deliverToFinal(thisLetter, finalIpsum)
		*current++
	}
}

func main() {

	runtime.GOMAXPROCS(2)

	index := new(int)
	*index = 0
	loremIpsum = "Lorem ipsum dolor sit amet consectetur, adipisicing elit. Expedita ducimus iste quas iusto deleniti et, fugit molestias voluptatem quisquam accusamus animi eaque ea, necessitatibus cumque possimus eligendi beatae, itaque suscipit."

	letters := []byte(loremIpsum)
	length := len(letters)

	go capitalize(index, length, letters, &finalIpsum)
	go func() {
		go capitalize(index, length, letters, &finalIpsum)
	}()

	fmt.Println(length, " characters.")
	fmt.Println(loremIpsum)
	fmt.Println(*index)
	fmt.Println(finalIpsum)
}
