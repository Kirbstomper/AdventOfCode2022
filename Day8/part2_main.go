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
			pos := checkIfVisibleUp(trees, r, c) * checkIfVisibleDown(trees, r, c) * checkIfVisibleLeft(trees, r, c) * checkIfVisibleRight(trees, r, c)
			if pos > answer {
				answer = pos
			}
		}
	}

	//Calculate the answer

	println(answer)
}

// returns true if tree is visible
func checkIfVisibleUp(trees [][]int, x int, y int) int {
	trees_vis := 0
	for i := y - 1; i >= 0; i-- {
		if trees[x][i] >= trees[x][y] {
			trees_vis++
			break
		} else {
			trees_vis++
		}
	}
	return trees_vis
}

func checkIfVisibleDown(trees [][]int, x int, y int) int {
	trees_vis := 0
	for i := y + 1; i < len(trees[x]); i++ {
		if trees[x][i] >= trees[x][y] {
			trees_vis++
			break
		} else {
			trees_vis++
		}
	}
	return trees_vis
}

func checkIfVisibleLeft(trees [][]int, x int, y int) int {
	trees_vis := 0
	for i := x - 1; i >= 0; i-- {
		if trees[i][y] >= trees[x][y] {
			trees_vis++
			break
		} else {
			trees_vis++
		}
	}
	return trees_vis
}

func checkIfVisibleRight(trees [][]int, x int, y int) int {
	trees_vis := 0
	for i := x + 1; i < len(trees); i++ {
		if trees[i][y] >= trees[x][y] {
			trees_vis++
			break
		} else {
			trees_vis++
		}
	}
	return trees_vis
}
