package main

import "fmt"

func main() {
	youtubesubscribers := map[string] int{
		"Tutorials" : 789,
		"Comedy" :2456,
		"Family" :7890,

	}
    fmt.Println(youtubesubscribers["Tutorials"])
	fmt.Println(youtubesubscribers["Comedy"])
	fmt.Println(youtubesubscribers["Family"])
}
