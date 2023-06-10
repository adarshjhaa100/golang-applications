package aoc

import (
	"fmt"
	"os"
	"strconv"
	// "io"
)

func que1part1(){

	// inputPath := "./aoc/inputs/que2021_1_1.txt"
	inputPath := "./aoc/inputs/adventofcode.com_2021_day_1_input.txt"
	fmt.Println("AOC 2021 Question 1 part1 with input", inputPath)

	fpr, err := os.OpenFile(inputPath,os.O_RDONLY,0644)
	checkfile(err)
	defer fpr.Close()

	lines := readLinesBufio(fpr) 

	// for newLine := readLineScratch(fpr); newLine != nil ; newLine = readLineScratch(fpr){
		
	// 	lineValStr := string(newLine)
	
	// 	fmt.Println( lineValStr )

	// 	// strconv is used for conversion to/from string. Contains methods like 
	// 	// parseFloat, atoi, itoa etc.
	// 	intVal, err := strconv.Atoi(lineValStr) 
	// 	checkfile(err)

	// }

	depth := -1 
	timesDepthIncrease := 0
	for _, val := range lines {
		depthNew, err := strconv.Atoi(val)
		checkfile(err)

		if depth < depthNew && depth != -1{
			timesDepthIncrease++
		}

		depth = depthNew
	}

	fmt.Printf("Times depth incerase %v\n", timesDepthIncrease)
}


func que1part2(){
	inputPath := "./aoc/inputs/adventofcode.com_2021_day_1_input.txt"
	fmt.Println("AOC 2021 Question 1 part2 with input", inputPath)

	fpr, err := os.OpenFile(inputPath,os.O_RDONLY,0644)
	checkfile(err)
	defer fpr.Close()

	lines := readLinesBufio(fpr) 
	if(len(lines) <= 3){
		return 
	}

	// Function closure (here the function below is able to modify the depths slice)
	var depths []int
	func(){
		for _, val := range lines {
			val, err := strconv.Atoi(val)
			checkfile(err)
			depths = append(depths, val)
		}
	}()

	startIndex, endIndex := 0, 2
	sumWindow := depths[0] + depths[1] + depths[2]
	timesWindowInc := 0

	for i := 3; i < len(depths); i++ {

		fmt.Printf("StartIndex: %v, EndIndex: %v, currentSum: %v\n", 
					startIndex, endIndex, sumWindow)

		newSumWindow := sumWindow + depths[endIndex+1] - depths[startIndex]

		fmt.Println(newSumWindow)

		startIndex, endIndex = startIndex + 1, endIndex + 1

		if(newSumWindow > sumWindow){
			timesWindowInc++
		}
		sumWindow = newSumWindow
	}

	fmt.Printf("Times Window Size Inc. %#v\n", timesWindowInc)

}