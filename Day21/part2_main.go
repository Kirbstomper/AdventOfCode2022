package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	expression string
	value      int
}

func main() {

	answer := 3352886133840
	f, err := os.Open("elfinput")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	monk_map_main := make(map[string]*monkey)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":") // Split line

		val, err := strconv.Atoi(strings.Trim(line[1], " "))
		if err != nil {
			monk_map_main[line[0]] = &monkey{expression: line[1], value: 0}
		} else {
			monk_map_main[line[0]] = &monkey{value: val, expression: ""}
		}
	}
	for true {
		answer = answer - 1
	//	println(answer)

		monk_map := make(map[string]*monkey)
		for k, v := range monk_map_main {
			monk_map[k] = &monkey{v.expression, v.value}
		}

		monk_map["humn"].value = answer
		for monk_map["root"].expression != "" {

			for k, v := range monk_map {
				//fmt.Println(k, v.expression, v.value)

				//if there is an expression still
				if v.expression != "" {
					//Check if children have values
					exp := strings.Split(v.expression, " ")
					can_eval := monk_map[exp[1]].expression == "" && monk_map[exp[3]].expression == ""
					if can_eval {
						if k == "root" {
							//fmt.Println("answer:", monk_map["humn"].value)
						//	fmt.Println("tar:", monk_map[exp[3]].value)
						//	fmt.Println("cur:", monk_map[exp[1]].value)

							if monk_map[exp[1]].value == monk_map[exp[3]].value {
								fmt.Println("ANSWER:", monk_map["humn"])
							}
						}

						if exp[2] == "+" {
							v.value = monk_map[exp[1]].value + monk_map[exp[3]].value
						}
						if exp[2] == "-" {
							v.value = monk_map[exp[1]].value - monk_map[exp[3]].value

						}
						if exp[2] == "*" {
							v.value = monk_map[exp[1]].value * monk_map[exp[3]].value

						}
						if exp[2] == "/" {
							v.value = monk_map[exp[1]].value / monk_map[exp[3]].value
						}

						// fmt.Println(v)
						// fmt.Println(exp[1])
						// fmt.Println(exp[3])

						// fmt.Println("---------------------------")
						monk_map[k] = &monkey{"", v.value}
					}
				}
			}

			//fmt.Println(monk_map["root"].value)

		}
	}
}
