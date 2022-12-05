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
	//Strings are slices right, so lets just use those as the stack
	//Also screw trying to parse the stack, so manually doing here
	stack := make([]string, 9)

	stack[0] = "LNWTD"
	stack[1] = "CPH"
	stack[2] = "WPHNDGMJ"
	stack[3] = "CWSNTQL"
	stack[4] = "PHCN"
	stack[5] = "THNDMWQB"
	stack[6] = "MBRJGSL"
	stack[7] = "ZNWGVBRT"
	stack[8] = "WGDNPL"

	for scanner.Scan() {

		line := scanner.Text()
		var split = strings.Split(line, " ")
		var num_move, _ = strconv.Atoi(split[1])
		var fr_stack, _ = strconv.Atoi(split[3])
		var to_stack, _ = strconv.Atoi(split[5])

		fr_stack--
		to_stack--
		fmt.Println(num_move, " ", fr_stack, " ", to_stack)

		//Now just do the actual moving
		for i := 0; i < num_move; i++ {
			//Pop from fr_stack to to_stack
			stack[to_stack] += string(stack[fr_stack][len(stack[fr_stack])-1])
			stack[fr_stack] = stack[fr_stack][:len(stack[fr_stack])-1]
		}
	}

	for _, s := range stack {
		println(s)
	}
}
