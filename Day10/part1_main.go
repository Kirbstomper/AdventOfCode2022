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
		cycles[i+1] += running
		fmt.Println(i+1, " : ", cycles[i])

	}
	fmt.Println((cycles[19] * 20) + (cycles[59] * 60) + (cycles[99] * 100) + (cycles[139] * 140) + (cycles[179] * 180) + (cycles[219] * 220))
}
