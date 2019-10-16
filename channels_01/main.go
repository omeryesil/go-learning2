//Check health status of different sites
//	Go scheduler handles the go routines
// 	As default, go scheduler uses only one CPU even with multi core CPUs
//	However, we can enable multi core cpus, in this case the Go scheduler assignes different go routines to different CPUs
//
//NOTE: Concurrency is not parallelism
//			Concurrency: We can have multiple threads executing the code, if one thread blocks, another one is picked up...
//			Parallelism: Multiple threads are executed at the same time, requires multiple CPUs
//
//Channels : Allows different go routines to communicate. Channels are the only way provides it.
//		Sending data with channels. Different ways:
//			channel <- 5 			//Send value of 5 to channel
//			myNumber <- channel 	//Wait for a value to be send into the channel. When we get one, then assign the value to myNumber
//			fmt.Println(<-channel)	//Wait for a value to be sent into the channel. When we get one, then log it out immediately.
//
//Literals : Same as Lambda in C# / Java
//
//
//

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://google.com",
		"https://msn.com",
		"https://facebook.com",
		"https://amazon.com",
	}

	//Creating a new channel, which will have string value.
	// 	Remember, make will create a new value out of the given type (in this case it is string)
	c := make(chan string)

	for _, link := range links {
		//with go command we can run the code in a separete go routine (like multi threads)
		go checkLink(link, c)
	}

	////Syntax 1: for infinite loop of checking the site statuses
	//for {
	//	go checkLink(<-c, c)
	//}

	//Syntax 2: for infinite loop of checking the site statuses
	for l := range c {
		//time.Sleep(time.Second) //This is not a right place to put the sleep, as time.Sleep only sleeps the current go routine.
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, chnl chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
	} else {
		fmt.Println(link, "is up")
	}

	chnl <- link
}
