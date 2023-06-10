package aoc

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)



func que2part1() {
	fp, err := os.Open("./aoc/inputs/adventofcode.com_2021_day_2_input.txt")
	checkfile(err)
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	horizontalPos, verticalPos := 0,0
	movementType, movementLength := "", 0

	//sc.Scan returns true until next token is read
	for sc.Scan(){
		line := sc.Text()
		movementType = strings.Split(line, " ")[0]
		movementLength, err =  strconv.Atoi(strings.Split(line, " ")[1])
		checkfile(err)

		switch movementType  {
			case "forward":
				horizontalPos += movementLength
			case "down":
				verticalPos += movementLength
			case "up":
				if( verticalPos - movementLength >= 0 ){
					verticalPos -= movementLength
				}		
			default:
				panic("incorrect movement type: "+movementType)
		}
		

		
	}
	fmt.Printf("Vertical Posn: %#v, Horizontal Pos: %#v, ans: %v \n", verticalPos, horizontalPos, verticalPos * horizontalPos )

}

func que2part2() {
	fp, err := os.Open("./aoc/inputs/adventofcode.com_2021_day_2_input.txt")
	checkfile(err)
	defer fp.Close()

	sc := bufio.NewScanner(fp)

	horizontalPos, verticalPos, aim := 0,0,0
	movementType, movementLength := "", 0

	//sc.Scan returns true until next token is read
	for sc.Scan(){
		line := sc.Text()
		movementType = strings.Split(line, " ")[0]
		movementLength, err =  strconv.Atoi(strings.Split(line, " ")[1])
		checkfile(err)

		switch movementType  {
			case "forward":
				horizontalPos += movementLength
				if( verticalPos + aim * movementLength >= 0 ){
					verticalPos = verticalPos + aim * movementLength
				}
			case "down":
				aim += movementLength
			case "up":
				aim -= movementLength		
			default:
				panic("incorrect movement type: "+movementType)
		}
		

		
	}
	fmt.Printf("Vertical Posn: %#v, Horizontal Pos: %#v, ans: %v \n", verticalPos, horizontalPos, verticalPos * horizontalPos )

}