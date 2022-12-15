package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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
	max_pos := 4000000
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",") // Split line
		sx, _ := strconv.Atoi(line[0])
		sy, _ := strconv.Atoi(line[1])
		bx, _ := strconv.Atoi(line[2])
		by, _ := strconv.Atoi(line[3])
		beacons[pos{x: bx, y: by}] = true
		//Get manhattan distance to closest sensor
		manhat := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		manHatDistances[pos{x: sx, y: sy}] = manhat
		//Fill in sensor range for each combination
	}

	//coveredPos := make(map[pos]bool)

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		findAnswer(0, 1000000, max_pos, manHatDistances)
		defer wg.Done()
	}()

	wg.Add(1)

	go func() {
		findAnswer(1000000, 2000000, max_pos, manHatDistances)
		defer wg.Done()
	}()

	wg.Add(1)

	go func() {
		findAnswer(2000000, 3000000, max_pos, manHatDistances)
		defer wg.Done()
	}()

	wg.Add(1)

	go func() {
		findAnswer(3000000, 4000000, max_pos, manHatDistances)
		defer wg.Done()
	}()
	wg.Wait()

	//fmt.Println(len(coveredPos))
}

func findAnswer(start, end, max_pos int, manHatDistances map[pos]int) {
	fmt.Println("Starting x from ", start, " to ", end)
	for x := start; x <= end; x++ {
		for y := 0; y <= max_pos; y++ {
			isCovered := false
			for k, v := range manHatDistances {
				if !isCovered {
					isCovered = k.isCovered(x, y, v)
				}
			}
			if !isCovered {
				fmt.Println("x:", x, "y:", y)
				fmt.Println(x*4000000 + y)
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
