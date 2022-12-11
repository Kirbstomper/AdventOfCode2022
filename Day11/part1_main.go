package main

import "fmt"

func main() {

	monkey_items := [][]int{{89, 73, 66, 57, 64, 80},
		{83, 78, 81, 55, 81, 59, 69},
		{76, 91, 58, 85},
		{71, 72, 74, 76, 68},
		{98, 85, 84},
		{78},
		{86, 70, 60, 88, 88, 78, 74, 83},
		{81, 58},
	}

	monkey_operation := make([]string, 8)
	monkey_operation[0] = "M"
	monkey_operation[1] = "A"
	monkey_operation[2] = "M"
	monkey_operation[3] = "M"
	monkey_operation[4] = "A"
	monkey_operation[5] = "A"
	monkey_operation[6] = "A"
	monkey_operation[7] = "A"

	monkey_mod := make([]int, 8)
	monkey_mod[0] = 3
	monkey_mod[1] = 1
	monkey_mod[2] = 13
	monkey_mod[3] = 0
	monkey_mod[4] = 7
	monkey_mod[5] = 8
	monkey_mod[6] = 4
	monkey_mod[7] = 5

	monkey_test := make([]int, 8)
	monkey_test[0] = 13
	monkey_test[1] = 3
	monkey_test[2] = 7
	monkey_test[3] = 2
	monkey_test[4] = 19
	monkey_test[5] = 5
	monkey_test[6] = 11
	monkey_test[7] = 17

	monkey_throw := [][]int{
		{6, 2},
		{7, 4},
		{1, 4},
		{6, 0},
		{5, 7},
		{3, 0},
		{1, 2},
		{3, 5},
	}

	insp_count := make([]int, 8)
	//Simulate monkey rounds

	for r := 0; r < 20; r++ { //For 20 rounds
		fmt.Println(r, " : ", monkey_items)
		for m := 0; m < len(monkey_items); m++ { //for the 8 monkeys..
			//Inspect each item
			fmt.Println("Monkey: ", m)
			for pos, i := range monkey_items[m] {

				fmt.Println("inspects item with level: ", i)
				insp_count[m]++ //monkey has inspected an item
				modifier := monkey_mod[m]
				//Modifiy
				if modifier == 0 {
					modifier = i
				}
				if monkey_operation[m] == "A" {
					monkey_items[m][pos] = monkey_items[m][pos] + modifier
					fmt.Println("Worry icrease by", modifier, "to ", monkey_items[m][pos])

				} else {
					monkey_items[m][pos] = monkey_items[m][pos] * modifier
					fmt.Println("Worry multiplied by", modifier, "to ", monkey_items[m][pos])
				}

				//Divide by 3
				monkey_items[m][pos] = monkey_items[m][pos] / 3
				fmt.Println("Monkey bored, worry level divided by 3 to ", monkey_items[m][pos])
				//Test
				if monkey_items[m][pos]%monkey_test[m] == 0 {
					//give them to the next monkey
					monkey_items[monkey_throw[m][0]] = append(monkey_items[monkey_throw[m][0]], monkey_items[m][pos])
				} else {
					monkey_items[monkey_throw[m][1]] = append(monkey_items[monkey_throw[m][1]], monkey_items[m][pos])
				}
			}
			monkey_items[m] = []int{}
		}
	}
	fmt.Println(insp_count)
}
