package main

import "github.com/kazukousen/cinema/scrape/scrape"

func main() {
	base := "https://eiga.com"
	path := "/theater/"
	selector := ""
	scrape.Run(base, path, selector)
}
