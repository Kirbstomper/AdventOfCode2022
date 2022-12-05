package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("elfInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	var overlap_sum = 0
	for scanner.Scan() {
		line := scanner.Text()

		var fst_rng, sec_rng = strings.Split(strings.Split(line, ",")[0], "-"), strings.Split(strings.Split(line, ",")[1], "-")
		var has_overlap = false
		first := make([]int, len(fst_rng))
		second := make([]int, len(sec_rng))
		for i, s := range fst_rng {
			first[i], _ = strconv.Atoi(s)
		}
		for i, s := range sec_rng {
			second[i], _ = strconv.Atoi(s)
		}

		fmt.Println(first[0], "-", first[1])
		fmt.Println(second[0], "-", second[1])

		if first[0] <= second[0] {
			if first[1] >= second[0] {
				println("FIRST LIST HAS IT")
				has_overlap = true
			}
		}
		if second[0] <= first[0] {
			if second[1] >= first[0] {
				println("SECOND LIST HAS IT")
				has_overlap = true
			}
		}
		if has_overlap {
			overlap_sum++
		}
	}

	fmt.Println("The Answer is: ", overlap_sum)

}
