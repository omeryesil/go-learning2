package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	arguments := os.Args[1:]

	if len(arguments) == 0 {
		log.Fatal("File name must be provided")
		os.Exit(-1)
	}

	option1(arguments[0])

	option2(arguments[0])

}

//uses interface
func option1(fileName string) {

	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal("An error has occured while reading the file: ", err)
		os.Exit(-1)
	}

	io.Copy(os.Stdout, f)
}

//Does not use interface struct
func option2(fileName string) {

	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatal("An error has occured while reading the file: ", err)
		os.Exit(-1)
	}
	s := string(bytes)

	fmt.Println(s)
}
