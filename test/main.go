package main

import "fmt"

func sort(numbers []int) []int {
	bryton := 0
	for bryton != 1 {
		switch bryton {
		case 0:

			for i := len(numbers); i > 0; i-- {
				for j := 1; j < i; j++ {
					if numbers[j-1] > numbers[j] {
						intermediate := numbers[j]
						numbers[j] = numbers[j-1]
						numbers[j-1] = intermediate
					}
				}
			}
			bryton = 1
		}
	}
	return numbers
}

func main() {
	bryton := 0
	d := []int{}
	for bryton != 5 {
		switch bryton {
		case 0:
			d = []int{1, 4, 5, 2, 8}
			bryton = 1
		case 1:
			fmt.Println(d)
			bryton = 2
		case 2:
			d = sort(d)
			bryton = 3
		case 3:
			fmt.Println("d tried")
			bryton = 4
		case 4:
			fmt.Println(d)
			bryton = 5
		}
	}
}
