package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	visited := make(map[pos]bool)
	//Read all data to a 2d array

	head_pos := pos{0, 0}
	tail_pos := pos{0, 0}

	visited[tail_pos] = true
	for scanner.Scan() {
		line := scanner.Text()

		inst := strings.Split(line, " ")
		amount_moved, _ := strconv.Atoi(inst[1])

		for amount_moved > 0 {
			if "R" == inst[0] {
				head_pos.x++
			}
			if "L" == inst[0] {
				head_pos.x--
			}
			if "U" == inst[0] {
				head_pos.y++
			}
			if "D" == inst[0] {
				head_pos.y--
			}
			fmt.Println("head: ", head_pos)

			//Calculate where the tail should be

			//IF they are touching still, no move needed for tail
			//Touching is left/right, up/down, diag, or overlapping
			if !isTouching(tail_pos, head_pos) {
				if "R" == inst[0] {
					tail_pos.x = head_pos.x - 1
					tail_pos.y = head_pos.y
				}
				if "L" == inst[0] {
					tail_pos.x = head_pos.x + 1
					tail_pos.y = head_pos.y

				}
				if "U" == inst[0] {
					tail_pos.x = head_pos.x
					tail_pos.y = head_pos.y - 1
				}
				if "D" == inst[0] {
					tail_pos.x = head_pos.x
					tail_pos.y = head_pos.y + 1
				}
				fmt.Println("tail: ", tail_pos)
				visited[tail_pos] = true
			}

			amount_moved--
		}

	}

	//Calculate the answer
	answer := 0
	for _, v := range visited {
		if v {
			answer++
		}

	}
	fmt.Println("the answer is: ", answer)
}

type pos struct {
	x, y int
}

// returns true if tree is visible
func isTouching(tail pos, head pos) bool {

	if math.Abs(float64(tail.x)-float64(head.x)) > 1 {
		return false
	}
	if math.Abs(float64(tail.y)-float64(head.y)) > 1 {
		return false
	}
	return true
}
