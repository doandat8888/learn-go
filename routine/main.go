package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello world")
	links := []string{
		"https://google.com",
		"https://amazon.com",
		"https://youtube.com",
	}

	//Create channel:
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error: ", err)
		c <- link
		return
	}
	fmt.Println("Successfully access to", link)
	c <- link
}
