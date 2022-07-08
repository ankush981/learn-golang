package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is a text line in a simple file"
	filename := "./myfile.txt"

	file, err := os.Create(filename)
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Length is:", length)
	defer file.Close()

	ReadFile(filename)
}

func ReadFile(filename string) {
	content, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("File contents:", content)
	fmt.Println("File contents as string:", string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
