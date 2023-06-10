package aoc

import "fmt"

const ENDLINE = '\n'

func SolveQuestion2021(questionCode string){
	
	switch questionCode {
		case "2021_1":
			que1part1()
			que1part2()
		case "2021_2":
			que2part1()
			que2part2()
	
		default:
			fmt.Println("Incorrect question code entered")
	
	}
}