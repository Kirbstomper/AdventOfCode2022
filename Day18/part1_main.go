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
	f, err := os.Open("elfinput")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	blocks := make([][100][100]bool, 100)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",") // Split line

		x, _ := strconv.Atoi(line[0])
		y, _ := strconv.Atoi(line[1])
		z, _ := strconv.Atoi(line[2])

		blocks[x+1][y+1][z+1] = true

	}

	//Find answer

	ans := 0
	for x := range blocks {
		for y := range blocks[x] {
			for z := range blocks[x][y] {
				if blocks[x][y][z] == true {
					fmt.Println(x, y, z)
					sides := 6
					//block to the left
					if blocks[x-1][y][z] {
						sides--
					}
					//block to the right
					if blocks[x+1][y][z] {
						sides--
					}
					//block in front
					if blocks[x][y-1][z] {
						sides--
					}
					//block behind
					if blocks[x][y+1][z] {
						sides--
					}
					//block on top
					if blocks[x][y][z-1] {
						sides--
					}
					//block below
					if blocks[x][y][z+1] {
						sides--
					}
					ans += sides
				}
			}
		}
	}

	fmt.Println("Answer is: ", ans)

}

type pos struct {
	x, y int
}
