package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	//Variable to hold the current largest # of calories
	var largest_cals_sum = 0

	f, err := os.Open("elfInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	list := make([]int, 0)

	var current_cals = 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			list = append(list, current_cals) //Append to a slice
			current_cals = 0
		} else {
			cals, _ := strconv.Atoi(line)
			current_cals += cals
		}
	}

	sort.IntSlice.Sort(list)  //Sort the slice
	list = list[len(list)-3:] //get the last 3 elements of the sorted slice
	largest_cals_sum = list[0] + list[1] + list[2]

	fmt.Println("The Answer is: ", largest_cals_sum)

}
