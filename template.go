package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("elfinput")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ") // Split line
	}
}

type pos struct {
	x, y int
}
