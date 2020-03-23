package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://amazon.in",
		"https://bing.com",
		"https://testme.com",
	}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		// function literal in GO
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)

	}
}

func checkLink(link string, c chan string) {

	time.Sleep(2 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("The link ", link, " is down ! ")
		c <- link
		return
	}
	fmt.Println("Link ", link, "seems working fine")
	c <- link
}
