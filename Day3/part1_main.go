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
	var priorities = "_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var priority_sum = 0
	for scanner.Scan() {
		line := scanner.Text()
		//Split line into 2
		var frs_half = string(line[:len(line)/2])
		var sec_half = string(line[len(line)/2:])

		println(frs_half)
		println(sec_half)

		//Count chars in first half
		var rune_map = make(map[rune]int)
		var repeat = ""
		for _, s := range frs_half {
			rune_map[s]++
		}

		//Check second second half for the repeat
		for _, s := range sec_half {
			if rune_map[s] > 0 {
				repeat = string(s)
				break
			}
		}
		println(repeat)
		priority_sum += strings.Index(priorities, repeat)
	}

	fmt.Println("The Answer is: ", priority_sum)

}
