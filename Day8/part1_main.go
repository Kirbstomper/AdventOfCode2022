package main

import (
	"bufio"
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

	trees := make([][]int, 0)
	//Read all data to a 2d array

	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, v := range strings.Split(line, "") {
			tree, _ := strconv.Atoi(v)
			row = append(row, tree)
		}
		trees = append(trees, row)
	}
	//Go over each tree 1 by 1 to check each direction

	answer := 0
	for r, _ := range trees {
		for c, _ := range trees[r] {

			//check up
			if checkIfVisibleUp(trees, r, c) || checkIfVisibleDown(trees, r, c) || checkIfVisibleLeft(trees, r, c) || checkIfVisibleRight(trees, r, c) {
				answer++
			}
			//check down
			//Check left
			//Check right
		}
	}

	//Calculate the answer

	println(answer)
}

// returns true if tree is visible
func checkIfVisibleUp(trees [][]int, x int, y int) bool {
	can_see := true
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= trees[x][y] {
			can_see = false
		}
	}
	return can_see
}

func checkIfVisibleDown(trees [][]int, x int, y int) bool {
	can_see := true
	for i := y + 1; i < len(trees[x]); i++ {
		if trees[x][i] >= trees[x][y] {
			can_see = false
		}
	}
	return can_see
}

func checkIfVisibleLeft(trees [][]int, x int, y int) bool {
	can_see := true
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= trees[x][y] {
			can_see = false
		}
	}
	return can_see
}

func checkIfVisibleRight(trees [][]int, x int, y int) bool {
	can_see := true
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= trees[x][y] {
			can_see = false
		}
	}
	return can_see
}
