package main

func main() {
	trends := Scrape()
	Send(trends)
}
