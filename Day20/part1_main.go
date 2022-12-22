package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("elfinput")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	order_list := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		val, _ := strconv.Atoi(line)
		order_list = append(order_list, val)
	}

	list_ring := ring.New(len(order_list))
	positions := make(map[element]*ring.Ring, len(order_list))

	//initialize ring
	var zero element
	for ind, val := range order_list {

		//fmt.Println(ind, val)
		toAdd := element{index: ind, value: val}
		if val == 0 {
			zero = toAdd
		}
		positions[toAdd] = list_ring
		list_ring.Value = val
		list_ring = list_ring.Next()
	}

	for ind, val := range order_list {
		r := positions[element{index: ind, value: val}].Prev()
		remove := r.Unlink(1)
		r.Move(val).Link(remove)
	}

	ans := 0
	list_ring = positions[zero]
	for i := 1; i <= 3; i++ {
		list_ring = list_ring.Move(1000)
		ans += list_ring.Value.(int)
	}

	fmt.Println(ans)
}

type element struct {
	index, value int
}
