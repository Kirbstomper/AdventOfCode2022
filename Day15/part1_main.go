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
	line_care := 2000000 // The line we wish to know how many spaces are covered on
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

	beacons_on_line := 0
	for k, _ := range beacons {
		if k.y == line_care {
			beacons_on_line++
		}
	}

	coveredPos := make(map[pos]bool)
	for i := -5000000; i < 5000000; i++ {
		for k, v := range manHatDistances {
			if k.isCovered(i, line_care, v) {
				coveredPos[pos{x: i, y: line_care}] = true
			}
		}
	}

	fmt.Println(len(coveredPos) - beacons_on_line)
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
