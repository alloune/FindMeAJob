package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Scraper Service Started")
	// Keep the service running
	for {
		time.Sleep(1 * time.Hour)
	}
}
