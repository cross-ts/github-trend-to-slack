package main

import "fmt"

func main() {
	trends := Scrape()
	fmt.Println(trends)
	Send()
}
