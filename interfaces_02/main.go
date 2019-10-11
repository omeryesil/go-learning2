package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const link string = "https://google.ca"

type logWriter struct{}

func main() {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(-1)
	}

	// LONG WAY TO GET response body -----------------------------------------------
	////creates an empty byte slice with the length of 99999
	////	this is not a real life scenario
	//bs := make([]byte, 99999)
	////Body.Read is an interface
	//resp.Body.Read(bs)
	//fmt.Println(string(bs))
	//io.Copy(os.Stdout, resp.Body)

	// WITH Copy with Writer Interface
	// io.Copy is : func Copy(dst Writer, src Reader) (written int64, err error)
	//		Writer is : type Writer interface {  Write(p []byte) (n int, err error)  }
	//		Reader is : type Reader interface {  Read(p []byte) (n int, err error)  }
	//io.Copy(os.Stdout, resp.Body)

	//NOT We are using our implementation of the interface
	fmt.Println("-------------------------------------------------------")
	lw := logWriter{}
	l, _ := io.Copy(lw, resp.Body)

	fmt.Println("Length", l)
}

//Here we are implementing the Writer interface in io package : type Writer interface { Write(p []byte) (n int, err error) }
func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
