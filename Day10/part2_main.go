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
	f, err := os.Open("elfinput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	//Read all data to a 2d array

	cycles := make(map[int]int)
	for i := 1; i < 300; i++ {
		cycles[i] = 1

	}
	stack := make([]int, 0)
	cycleMap := make(map[int]int, 0)

	for i := 1; i < 250; i++ {
		cycleMap[i] = 1
	}
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		if split[0] == "addx" {
			c, _ := strconv.Atoi(split[1])
			stack = append(stack, 0)
			stack = append(stack, c)
		} else {
			stack = append(stack, 0)
		}

	}

	running := 0
	for i := 0; i <= 300; i++ {
		if len(stack)-1 > i {
			running += stack[i]
		}
		cycles[i+2] += running
		fmt.Println(i+1, " : ", cycles[i])

	}
	fmt.Println((cycles[20] * 20) + (cycles[60] * 60) + (cycles[100] * 100) + (cycles[140] * 140) + (cycles[180] * 180) + (cycles[220] * 220))

	//LEts draw!

	for i := 0; i < 6; i++ {
		for r := 1; r <= 40; r++ {
			c := cycles[r+(i*40)]
			if c == r || c+1 == r || c+2 == r {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}
