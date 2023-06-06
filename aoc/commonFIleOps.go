package aoc

import (
	// "fmt"
	// "io"
	"bufio"
	"os"
	// "unicode"
	// "strings"
)

const EOL = "\n"
const NULL_CHAR = 0x0


func checkfile(err error){
	if(err != nil){
		panic(err)		
	}
}

// read and return a line from file using file pointer. This shifts the file pointer to the nextline if available 
func readLineScratch(fp *os.File) []byte{
	readByte := make([]byte, 1)
	var lineOut []byte
	err := error(nil)

	for ; (err==nil) && (string(readByte[0])!=EOL) ; _,err = fp.Read(readByte) {
		if(readByte[0] != NULL_CHAR){
			lineOut = append(lineOut, readByte...)
		}
	}
	return lineOut
}


// Reading files line by line using bufio package: 
// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
func readLinesBufio(fp *os.File) []string{
	var lines []string

	scanner := bufio.NewScanner(fp)

	// read the next token. By default the next token is new line
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // scanner.Text() returns a string
	}

	checkfile(scanner.Err())

	return lines	
}