package main

import "fmt"

func main(){
	phones := make(map[string]string)

	phones["huber"] = "0-800-huber"
	phones["huber"] = "0-800-huber2"
	phones["juan"] = "0-800-juan"
	phones["maria"] = "0-800-maria"
	for k, v := range phones{
		fmt.Printf("The phone of %s is: %s\n", k, v)
	}

	delete(phones, "juan")

	for k := range phones{
		fmt.Printf("Contact: %s\n", k)
	}

}
