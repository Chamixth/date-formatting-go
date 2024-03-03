package main

import (
	"fmt"
	"time"
)

func main() {
	// Given string
	t := time.Now()

	timePart := t.Format("15:04")
	fmt.Println("Time:", timePart)
	fmt.Println("Date: ")
	date()
}
func date() error {
	// Given string
	t := time.Now()

	 // Format the date as yyyy-mm-dd
	 formattedDate := t.Format("2006-01-02")

	 // Print the formatted date
	 fmt.Println("Formatted date:", formattedDate)
	
	return nil
}
