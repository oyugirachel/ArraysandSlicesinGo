package main

import "fmt"

var days [] string
func main() {



	days := [...]string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	weekdays := days[0:5]

	fmt.Println(days[0])
	fmt.Println(days[5])
	fmt.Println(weekdays)
}
