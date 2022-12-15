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
	f, err := os.Open("elfinput_real")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	manHatDistances := make(map[pos]int)
	beacons := make(map[pos]bool)
	pos_to_check := make(map[pos]int)
	max_pos := 4000000

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",") // Split line
		sx, _ := strconv.Atoi(line[0])
		sy, _ := strconv.Atoi(line[1])
		bx, _ := strconv.Atoi(line[2])
		by, _ := strconv.Atoi(line[3])
		beacons[pos{x: bx, y: by}] = true
		//Get manhatton distance to closest sensor
		manhat := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		manHatDistances[pos{x: sx, y: sy}] = manhat
		//Fill in sensor range for each combination

	}

	//build positions to check map
	for k, v := range manHatDistances {
		for i := 0; i < v; i++ { //Above left
			pos_to_check[pos{k.x - i, k.y - (v + 1) - i}] = 0
		}
		for i := 0; i < v; i++ { //Above right
			pos_to_check[pos{k.x + i, k.y - (v + 1) - i}] = 0
		}
		for i := 0; i < v; i++ { //bellow left
			pos_to_check[pos{k.x - i, k.y + (v + 1) - i}] = 0
		}
		for i := 0; i < v; i++ { //bellow right
			pos_to_check[pos{k.x + i, k.y + (v + 1) - i}] = 0
		}
	}

	println(len(pos_to_check))
	fmt.Println("Checking")
	///check if position is covered
	for p, _ := range pos_to_check {
		for k, v := range manHatDistances {
			if k.isCovered(p.x, p.y, v) {
				pos_to_check[p]++
			}
		}
	}

	fmt.Println("finding Answer")
	for p, v := range pos_to_check {
		if v == 0 {
			if 0 <= p.x && p.x <= max_pos && 0 <= p.y && p.y <= max_pos {
				fmt.Println((p.x * 4000000) + p.y) //Should be the answer
			}
		}
	}

}

func (p pos) isCovered(x, y, distance int) bool {

	if int(math.Abs(float64(p.x-x))+math.Abs(float64(p.y-y))) > distance {
		return false
	}
	return true
}

type pos struct {
	x, y int
}
