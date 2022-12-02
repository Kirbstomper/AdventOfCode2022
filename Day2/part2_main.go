package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var shape_map = make(map[string]int)
	f, err := os.Open("elfInput.txt")
	shape_map["A"] = 1
	shape_map["B"] = 2
	shape_map["C"] = 3
	shape_map["X"] = 1
	shape_map["Y"] = 2
	shape_map["Z"] = 3

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var points = 0
	for scanner.Scan() {
		line := scanner.Text()

		var game = strings.Split(line, " ")

		var play_a = shape_map[game[0]]
		var play_b = shape_map[game[1]]

		if play_b == 1 { //lose

			if play_a == 1 {
				points += 3
			} else {
				points += (play_a - 1)
			}

		} else if play_b == 2 { //draw
			points += 3
			points += play_a
		} else if play_b == 3 { //win
			points += 6
			if play_a == 3 {
				points += 1
			} else {
				points += (play_a + 1)
			}
		}

	}

	fmt.Println("The Answer is: ", points)

}
