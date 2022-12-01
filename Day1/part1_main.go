package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	//Variable to hold the current largest # of calories
	var largest_cals = 0

	f, err := os.Open("elfInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var current_cals = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if current_cals > largest_cals {
				largest_cals = current_cals
			}
			current_cals = 0
		} else {
			cals, _ := strconv.Atoi(line)
			current_cals += cals
		}
	}

	fmt.Println("The Answer is: ", largest_cals)

}
