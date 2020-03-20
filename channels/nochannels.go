package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"https://www.google.com",
		"https://amazon.in",
		"https://bing.com",
		"https://testmeagrar.com",
	}
	for _, link := range links {
		checkLink(link)
	}

}

func checkLink(link string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println("The link ", link, " is down ! ")
		return
	}
	fmt.Println("Link ", link, "seems working fine")
}
