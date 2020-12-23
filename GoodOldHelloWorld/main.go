package main

import "fmt"

func main() {
	x := 17.0
	time := 1.0
	for time <= 10 {
		x = x * 1.1
		fmt.Println(time, x)
		time = time + 1
	}
}
