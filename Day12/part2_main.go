package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var end pos

func main() {
	f, err := os.Open("elfinput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	line_num := 0
	start_pos := []pos{}
	hill_map := make(map[pos]int)
	adjacent := make(map[pos][]pos)

	for scanner.Scan() {
		line := scanner.Text()
		vals := []int{}

		for i, s := range strings.Split(line, "") {
			if s == "a" {
				start_pos = append(start_pos, pos{x: i, y: line_num})
			}
			if s == "E" {
				s = "z"
				end = pos{x: i, y: line_num}
			}
			hill_map[pos{x: i, y: line_num}] = int(s[0])

			vals = append(vals, int(s[0]))
			adjacent[pos{x: i, y: line_num}] = append(adjacent[pos{x: i, y: line_num}], pos{x: i + 1, y: line_num})
			adjacent[pos{x: i, y: line_num}] = append(adjacent[pos{x: i, y: line_num}], pos{x: i - 1, y: line_num})
			adjacent[pos{x: i, y: line_num}] = append(adjacent[pos{x: i, y: line_num}], pos{x: i, y: line_num + 1})
			adjacent[pos{x: i, y: line_num}] = append(adjacent[pos{x: i, y: line_num}], pos{x: i, y: line_num - 1})
		}

		line_num++
	}

	for _, p := range start_pos {
		fmt.Println(p)
	}

	true_ans := math.MaxInt
	for _, start := range start_pos {
		visited := make(map[pos]bool)
		parent := make(map[pos]pos)
		pos_queue := []pos{}
		//fmt.Println(end)

		//fmt.Println(hill_map)
		//Fuck it lets BFS
		visited[start] = true
		pos_queue = append(pos_queue, start)

		for len(pos_queue) > 0 {
			s := pos_queue[0]
			//fmt.Println(s, "value: ", hill_map[s], string(hill_map[s]))
			pos_queue = pos_queue[1:]

			if s.x == end.x {
				if s.y == end.y {
					//fmt.Println("WE DID IT")
					break
				}
			}

			for _, p := range adjacent[s] {
				//check if we can actually go to this neighbor
				if p.x >= 0 && p.y >= 0 && p.y < 41 && p.x < 78 {
					if (hill_map[s] >= hill_map[p]) || (hill_map[s] < hill_map[p] && math.Abs(float64(hill_map[s]-hill_map[p])) == 1) {
						if !visited[p] {
							//	fmt.Println("Adjacent is : ", p, " value is: ", hill_map[p], string(hill_map[p]))
							visited[p] = true
							parent[p] = s
							pos_queue = append(pos_queue, p)
						}
					}
				}
			}

		}
		ans := 0
		_, ok := parent[end]
		if ok {
			parent_end := end
			for parent_end != start {
				ans++
				parent_end = parent[parent_end]
			}
			if ans < true_ans {
				true_ans = ans
			}
			println(ans)
		}
	}

	println(true_ans)
}

type pos struct {
	x, y int
}
