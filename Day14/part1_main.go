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

	mount_map := make(map[pos]string)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ") // Split line

		positions := []pos{}
		for _, p := range line {
			pair := strings.Split(p, ",")
			x_pos, _ := strconv.Atoi(pair[0])
			y_pos, _ := strconv.Atoi(pair[1])
			positions = append(positions, pos{x: x_pos, y: y_pos})
		}
		start := positions[0]

		for i := 1; i < len(positions); i++ {
			end := positions[i]
			if start.x == end.x {
				if start.y > end.y {
					for r := start.y; r >= end.y; r-- {
						mount_map[pos{start.x, r}] = "#"
					}
				}
				if start.y < end.y {
					for r := start.y; r <= end.y; r++ {
						mount_map[pos{start.x, r}] = "#"
					}
				}
			}
			if start.y == end.y {
				if start.x > end.x {
					for r := start.x; r >= end.x; r-- {
						mount_map[pos{r, start.y}] = "#"
					}
				}
				if start.x < end.x {
					for r := start.x; r <= end.x; r++ {
						mount_map[pos{r, start.y}] = "#"
					}
				}
			}
			start = positions[i]
		}

	}
	printMap(mount_map)

	//Simulate sand
	answer := 0
	no_overflow := true
	for no_overflow {
		answer++
		sand_pos := pos{500, 0}
		phase := 0
		settled := false

		for !settled {
			if phase == 0 {
				if mount_map[pos{sand_pos.x, sand_pos.y + 1}] != "#" {
					sand_pos.y++
				} else {
					phase = 1
				}
			}
			if phase == 1 {
				if mount_map[pos{sand_pos.x - 1, sand_pos.y + 1}] != "#" {
					sand_pos.y++
					sand_pos.x--
					phase = 0
				} else {
					phase = 2
				}
			}
			if phase == 2 {
				if mount_map[pos{sand_pos.x + 1, sand_pos.y + 1}] != "#" {
					sand_pos.x++
					sand_pos.y++
				}
				phase = 0
			}

			if (mount_map[pos{sand_pos.x, sand_pos.y + 1}] == "#") && (mount_map[pos{sand_pos.x - 1, sand_pos.y + 1}] == "#") && (mount_map[pos{sand_pos.x + 1, sand_pos.y + 1}] == "#") {
				settled = true
				fmt.Println("SETTLED at ", sand_pos)
				mount_map[sand_pos] = "#"
			}
			if sand_pos.y > 10000 {
				fmt.Println(answer)
				printMap(mount_map)
				log.Panic()
			}
		}

	}
}

func printMap(pos_map map[pos]string) {

	visual := [510][15]string{}

	for k, _ := range pos_map {

		visual[k.x][k.y] = "#"
	}

	println()
	for y := 0; y <= 11; y++ {
		for x := 490; x < 510; x++ {
			if visual[x][y] != "#" {
				print(".")
			} else {
				print("#")
			}
		}
		println()
	}

}

type pos struct {
	x, y int
}
