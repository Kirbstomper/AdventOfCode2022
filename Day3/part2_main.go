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

		bag_1 := scanner.Text() //first bag
		scanner.Scan()
		bag_2 := scanner.Text() //second bag
		scanner.Scan()
		bag_3 := scanner.Text() //third bag

		//Count chars in first bag/
		var rune_map = make(map[rune]int)
		var repeat = ""
		for _, s := range bag_1 {
			if rune_map[s] == 0 {
				rune_map[s]++
			}
		}

		//Add to count in second bag if appears in first
		for _, s := range bag_2 {
			if rune_map[s] == 1 {
				rune_map[s]++
			}
		}
		//Check last bag. If count is 2 for an item it has appeared in the other two bags and is the repeat
		for _, s := range bag_3 {
			if rune_map[s] == 2 {
				repeat = string(s)
				break
			}
		}
		println(repeat)
		priority_sum += strings.Index(priorities, repeat)
	}

	fmt.Println("The Answer is: ", priority_sum)

}
