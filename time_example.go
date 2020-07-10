package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch t.Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend!")
	case time.Thursday:
		fmt.Println("Is Thursday of Bora Bora!")
	default:
		fmt.Println("Is a normal day!")
	}
}
