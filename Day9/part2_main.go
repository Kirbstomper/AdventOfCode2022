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

	knot_pos := make([]pos, 10)
	knot_pos_old := make([]pos, 10)

	visited[knot_pos[9]] = true
	for scanner.Scan() {
		line := scanner.Text()

		inst := strings.Split(line, " ")
		amount_moved, _ := strconv.Atoi(inst[1])

		for amount_moved > 0 {
			knot_pos_old[0] = knot_pos[0]
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

			//Calculate where the tail should be

			//IF they are touching still, no move needed for tail
			//Touching is left/right, up/down, diag, or overlapping
			dir := inst[0]
			for i := 1; i < 10; i++ {
				knot_pos_old[i] = knot_pos[i]

				head_pos := &knot_pos[i-1]
				tail_pos := &knot_pos[i]

				if !isTouching(*tail_pos, *head_pos) {
					if "R" == dir {
						tail_pos.x = head_pos.x - 1
						tail_pos.y = head_pos.y
					}
					if "L" == dir {
						tail_pos.x = head_pos.x + 1
						tail_pos.y = head_pos.y

					}
					if "U" == dir {
						tail_pos.x = head_pos.x
						tail_pos.y = head_pos.y - 1
					}
					if "D" == dir {
						tail_pos.x = head_pos.x
						tail_pos.y = head_pos.y + 1
					}
					if "DAG" == dir {
						knot_pos[i] = knot_pos_old[i-1]
					}

					dir = 
				}

			}

			visited[knot_pos[9]] = true
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
