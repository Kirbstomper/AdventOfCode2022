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

	knots_num := 10
	knot_pos := make([]pos, knots_num)

	visited[knot_pos[knots_num-1]] = true
	for scanner.Scan() {
		line := scanner.Text()

		inst := strings.Split(line, " ")
		amount_moved, _ := strconv.Atoi(inst[1])

		for amount_moved > 0 {

			//Move the head
			if "R" == inst[0] {
				knot_pos[0].x++
			}
			if "L" == inst[0] {
				knot_pos[0].x--
			}
			if "U" == inst[0] {
				knot_pos[0].y++
			}
			if "D" == inst[0] {
				knot_pos[0].y--
			}

			//IF they are touching still, no move needed for tail
			//Touching is left/right, up/down, diag, or overlapping
			for i := 1; i < knots_num; i++ {

				knot_pos[i] = getDirectionMoved(knot_pos[i], knot_pos[i-1])

			}
			//println("____-___________")
			visited[knot_pos[knots_num-1]] = true
			amount_moved--
		}
		for i, p := range knot_pos {
			fmt.Println(i, " : ", p)
		}
	}

	//Calculate the answer
	answer := 0
	for _, v := range visited {
		if v {
			//fmt.Println(k)
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
	if math.Abs(float64(tail.x)-float64(head.x)) > 1 || math.Abs(float64(tail.y)-float64(head.y)) > 1 {
		return false
	}
	return true
}

func getDirectionMoved(tail, head pos) pos {
	if isTouching(tail, head) {
		return tail
	}
	change_x := float64(head.x - tail.x)
	change_y := float64(head.y - tail.y)

	if math.Abs(change_x) > 1 || math.Abs(change_x)+math.Abs(change_y) > 2 {

		if change_x > 0 {
			tail.x++
		} else {
			tail.x--
		}

	}
	if math.Abs(change_y) > 1 || math.Abs(change_x)+math.Abs(change_y) > 2 {
		if change_y > 0 {
			tail.y++
		} else {
			tail.y--
		}
	}

	return tail
}
