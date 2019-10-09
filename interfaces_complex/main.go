package main

import (
	"fmt"
	"net/http"
	"os"
)

const link string = "https://google.ca"

func main() {
	resp, err := http.Get(link)

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(-1)
	}

	fmt.Println(resp.Body)

}
