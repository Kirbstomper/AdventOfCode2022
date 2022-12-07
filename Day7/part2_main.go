package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var allDirMap map[string]*Dir

func main() {
	allDirMap = make(map[string]*Dir)
	f, err := os.Open("elfInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	currentPath := ""
	//Map of all dirs and their size
	//Map of directories and total size for each
	// Key= DirName, val = total size
	//WHen .. is called, calculate total size for a directory

	//Read input and determine the full file size
	for scanner.Scan() {

		line := scanner.Text()

		command := strings.Split(line, " ")
	Start:
		if command[1] == "cd" {
			if command[2] == ".." {
				//chop off last part of current path
				splitPath := strings.Split(currentPath, "/")
				currentPath = strings.Join(splitPath[:len(splitPath)-2], "/") + "/"
			} else {
				currentPath += command[2] + "/"
				if _, ok := allDirMap[currentPath]; !ok {
					allDirMap[currentPath] = &Dir{SubDirs: make([]string, 0)}
				}
			}

		}

		if command[1] == "ls" {
			//Get all file sizes and dirs
			for scanner.Scan() {
				line = scanner.Text()
				command = strings.Split(line, " ")

				if command[0] == "$" {
					goto Start
				} else {
					if command[0] == "dir" {
						newSubDir := currentPath + command[1] + "/"
						allDirMap[currentPath].SubDirs = append(allDirMap[currentPath].SubDirs, newSubDir)
					} else {
						toAdd, _ := strconv.Atoi(command[0])
						allDirMap[currentPath].FileSize += toAdd
					}
				}
			}

		}

	}

	//print all dirs just for the hell of it

	answer := 100000000000
	//Get the size of all files on disk
	currentFree := 70000000 - allDirMap["root/"].getTotalSize()
	needToFree := 30000000 - currentFree
	for _, v := range allDirMap {
		amt := v.getTotalSize()
		if amt < answer {
			if amt >= needToFree {
				answer = amt
			}
		}
	}
	//Calculate the answer
	fmt.Println("The answer is: ", answer)

}

//Checks a string of characters for uniqueness

type Dir struct {
	SubDirs  []string
	FileSize int
}

// Returns the total size of a directory
func (d Dir) getTotalSize() int {
	totalSize := 0

	for _, sub := range d.SubDirs {
		totalSize += allDirMap[sub].getTotalSize()
	}

	return totalSize + d.FileSize
}
