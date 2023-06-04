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

	fmt.Println("File Reader...")

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


func fileWrite(){

	fmt.Println("File Writer...")

	// Dump full string into a file
	writeBuf := []byte("Hello, Im Writing to a File!")

	// The third argument is writeFile is permission flag. For more info: https://docs.nersc.gov/filesystems/unix-file-permissions/ 
	// The permission is generally represented in octal and is a 10 byte flag where: 0 -> directory; 1,2,3-> user; 4,5,6 -> group; 7,8,9 -> other(world)
	// e.g. Here we have 0644 which is in octal. 0 =  not a directory, 6 - rw for user, 4 , 4 - read for group and other 
	err := os.WriteFile("./data/testwritefull.txt", writeBuf, 0644)
	check(err)

	// For more granular wrtes open a file. Creates or Truncates(if already exists) the file
	fdw, err := os.Create("./data/fileWrite1.txt") // perm = 0666
	check(err)
	defer fdw.Close()

	// File system has an in memory copy. This function flushes that to stable storage
	fdw.Sync()

	// Write bytes to file using fs
	data2 := []byte{65,66,67}
	
	// write String into file
	data3 := "Writing String to a File"
	
	// fmt.Println(string(data2))
	n,err := fdw.Write(data2)
	check(err)
	fmt.Printf("Wrote %v bytes to file\n", n)

	// Write String
	n, err = fdw.WriteString(data3)
	check(err)
	fmt.Printf("Wrote %v bytes to file\n", n)


	// Write to an existing file. The second argument is a flag which can be or stacked O_APPEND|os.O_CREATE|os.O_WRONLY (this would append to file, if doesn't exist, create file). The constants can be obtained from: https://pkg.go.dev/os#pkg-constants
	fdw2, err := os.OpenFile("./data/fileWriter2.txt",os.O_CREATE|os.O_RDWR, 0644)
	check(err)
	defer fdw2.Close()

	bytesWrite := []byte("Hello World, Im appending to file\n")
	n, err = fdw2.Write(bytesWrite)
	check(err)
	fmt.Printf("Wrote %v bytes to file\n", n)


	fdw2.Seek(10,0)
	bytesWrite = []byte("Writing in the middle")
	n, err = fdw2.Write(bytesWrite)
	check(err)
	fmt.Printf("Wrote %v bytes to file\n", n)

}


func CallFileIOFunctions(){
	// fileRead()
	fileWrite()

}




