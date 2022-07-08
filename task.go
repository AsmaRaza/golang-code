package main

import (
	"fmt"
	"io/ioutil"
	//"strings"
)

//function for counting characters from textfile
func CountCharacters(data string) {
	c_count := 0
	for i := 0; i < len(data); i++ {
		c_count++
	}
	fmt.Println("The number of characters in a file =", c_count)

}

//function for counting words from text file
func CountWord(data string) {
	w_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			w_count++
		}
		// count = count + 1
	}
	fmt.Println("The number of  words in a file =", w_count)

}

//function for counting lines from textfile
func CountLines(data string) {
	l_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			l_count++
			//i++
		}

	}
	l_count++
	fmt.Println("The number of lines in a file =", l_count)

}

//function for counting whitespaces in a textfile

func CountWhiteSpaces(data string) {
	ws_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			ws_count++
		}

	}
	fmt.Println("The number of whitespaces in a file is =", ws_count)
}

func main() {
	data, err := ioutil.ReadFile("file.txt")
	//reading a file
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// fmt.Println("Contents of file:")
	// fmt.Println(string(data))
	CountCharacters(string(data))
	CountWord(string(data))
	CountLines(string(data))

	CountWhiteSpaces(string(data))

}
