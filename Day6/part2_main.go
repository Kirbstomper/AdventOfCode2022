package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("elfInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	//Strings are slices right, so lets just use those as the queue
	var answer = -1
	for scanner.Scan() {

		line := scanner.Text()
		var sub_buffer = ""

		for i, c := range line {
			sub_buffer += string(c)
			if len(sub_buffer) > 14 {
				//pop the first
				sub_buffer = sub_buffer[1:]
				//check unique
				if isUnique(sub_buffer) {
					answer = i + 1
					break
				}
			}
		}
	}
	fmt.Println("The answer is: ", answer)

}

// Checks a string of characters for uniqueness
func isUnique(buffer string) bool {
	for _, c := range buffer {
		if strings.Count(buffer, string(c)) > 1 {
			return false
		}
	}
	return true
}
