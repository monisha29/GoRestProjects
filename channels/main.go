package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//fmt.Println(<-c)
	// for {
	//	go checkLink(<-c, c)
	//}
	//Sleeping here throttle
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

//concurrency is different from parallelism
//channel are used to communicate between go routines
//Recieving messages froma achannel is a blocking call
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		//c <- "Might be down I think"
		c <- link
		return
	}

	fmt.Println(link, " is up !")
	//c <- "Yep its up"
	c <- link
}
