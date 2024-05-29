package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	size, err := f.Write([]byte("Writing data to file"))
	// size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("File created sucessfully! Size: %d bytes\n", size)
	f.Close()

	//reading
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	//reading little by little by opening the file
	file2, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file2)
	buffer := make([]byte, 3) //3 bytes
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
