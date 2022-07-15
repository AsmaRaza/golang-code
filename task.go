package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func runningtime(s string) (string, time.Time) {
	log.Println("Start:	", s)
	return s, time.Now()
}

func track(s string, startTime time.Time) {
	endTime := time.Now()
	log.Println("End:	", s, "took", endTime.Sub(startTime))
}

func execute() {
	defer track(runningtime("execute"))
	//time.Sleep(3 * time.Second)
}

//function for counting characters from textfile
func CountCharacters(data string, c chan string) {
	c_count := 0
	for i := 0; i < len(data); i++ {
		c_count++
	}

	result1 := fmt.Sprintln("The number of characters in a file is =", c_count)
	c <- result1

}

//function for counting words from text file
func CountWord(data string, c chan string) {
	w_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			w_count++

		}

	}

	result2 := fmt.Sprintln("The number of words in a file is =", w_count)
	c <- result2

}

//function for counting lines from textfile
func CountLines(data string, c chan string) {
	l_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '\n' {
			l_count++

		}

		result3 := fmt.Sprintln("The number of lines in a file is =", l_count)
		c <- result3
		l_count++

	}

}

//function for counting whitespaces in a textfile

func CountWhiteSpaces(data string, c chan string) {
	ws_count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == ' ' {
			ws_count++
		}

	}
	result4 := fmt.Sprintln("The number of whitespaces in a file is =", ws_count)
	c <- result4
}

func main() {
	c := make(chan string)

	execute()

	var path string

	flag.StringVar(&path, "path", "", "")
	flag.Parse()
	fmt.Println("Path", path)

	fmt.Println(path, "file.txt")

	data, err := ioutil.ReadFile("file.txt")
	//reading a file
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	// fmt.Println("Contents of file:")
	//fmt.Println(string(data))
	//go final(string(data))
	//w, z := <-c, <-c // receive from c

	go CountCharacters(string(data), c)
	go CountWord(string(data), c)
	go CountLines(string(data), c)

	go CountWhiteSpaces(string(data), c)
	for i := 0; i < 4; i++ {
		data := <-c
		fmt.Println(data)
	}
	// w, x := <-c, <-c
	// y, z := <-c, <-c
	// fmt.Println(w)
	// fmt.Println(x)
	// fmt.Println(y)
	// fmt.Println(z)

}
