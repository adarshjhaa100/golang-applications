package ioexamples

import (
	"io"
	"os"
	// "io"
	"bufio"
	"fmt"
	// "unsafe"
)


func check(e error){
	if e!=nil {
		panic(e)
	}
}

func fileRead(){

	// Get current working directory
	cwd, err := os.Getwd()
	check(err)
	fmt.Printf("current working directory: %v\n", cwd)

	// Read current directory
	directories, err := os.ReadDir(".")
	check(err)

	fmt.Println("List of files in the curent dir:")
	for _, dirEntry := range(directories){
		fmt.Printf("entry: %#v \n", dirEntry)
	}
 

	// Store the entire file contents into memory
	filePath := "./data/testread.txt"
	entireFile, err := os.ReadFile(filePath) 
	check(err)
	fmt.Printf("Size %#v\n", len(entireFile)) // the file contents by default are in binary format


	// Using file descriptor (type os.File)
	fd, err := os.Open(filePath)
	check(err)
	defer fd.Close()

	// check fileinfo
	fileInfo, err := fd.Stat()
	check(err)
	fmt.Printf("%#v \n", fileInfo)

	// read into filebuffer
	const BUFFER_SIZE = 5
	buff1 := make([]byte, BUFFER_SIZE)

	// Read upto the buffer size
	len, err := fd.Read(buff1)
	check(err)

	fmt.Printf("read %d bytes, %s\n", len, buff1[:len]) // only display upto bytes read
	

	// The read operation moves the filepointer by n bytes
	// fd.Seek(0, 1) can be used to check get the current offset (posn of current file ptr)
	offset, errOffset := fd.Seek(0, 1) // seek 0 from current posn(whence=1)
	check(errOffset)
	fmt.Printf("Offset after reading into buffer of 5 bytes for current file pointer: %v\n", offset)

	

	// fd.Seek(off, whence). seek right by 'off' bytes (if off>0 else seek left) and returns the new offset (posn)
	// whence is an indicator having 0,1,2 where 0 is to move from start, 1 is to move from current posn, 2 is from end 
	offset, errOffset = fd.Seek(6, 0) // move 6 bytes to right from start
	check(errOffset)
	buff2 := make([]byte, 5)
	len, err = fd.Read(buff2)
	check(err)

	fmt.Printf("Offset after fd.Seek(6,0) for current file pointer: %v\n", offset)
	fmt.Printf("Bytes read: %v, value: %#v from offset %v\n", len, string(buff2), offset)

	// Using additional methods from "io" library
	buff3 := make([]byte, 6)
	len, err = io.ReadAtLeast(fd, buff3, 6)
	check(err)
	offset, errOffset = fd.Seek(0, 1) // move 6 bytes to right from start
	check(errOffset)
	fmt.Printf("Bytes read: %v, byte value%#v: value: %#v from offset %#v\n", len, buff3, string(buff3), offset)
	

	// no builtin rewind, but can seek to start
	_, err = fd.Seek(0,0)
	check(err)


	// Reading using bufio: implements a buffered reader which is efficient for small reads
	// Also has some additional functions
	buffRdr := bufio.NewReader(fd)
	// buff4 := make([]byte, 10)// no need to explicitly create a new buffer
	bytesRead, errBuf := buffRdr.Peek(10) // Peek n bytes w/o moving the file reader
	check(errBuf)
	fmt.Printf("No of bytes read using bufio: %v, bytes: %v, str: %v ",10, bytesRead, string(bytesRead) )

}


func CallFileIOFunctions(){
	fileRead()
}




