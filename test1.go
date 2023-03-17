package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by 0")
	}
	return a / b, nil
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
func main() {
	fmt.Println("Hello, You")
	g := greeter{
		greeting: "Hello",
		name:     "GO",
	}
	g.greet()
	d, err := divide(5.0, 0.0)
	if err != nil {
	       fmt.Println(err)
	        return 
	}
	fmt.Println(d)
}
#https://go.dev/play/
#https://www.youtube.com/watch?v=YS4e4q9oBaU
#⭐️ Course Contents ⭐️
#⌨️ (0:00:00) Introduction 
#⌨️ (0:16:57) Setting Up a Development Environment
#⌨️ (0:35:48) Variables
#⌨️ (0:57:05) Primitives
#⌨️ (1:26:29) Constants
#⌨️ (1:47:53) Arrays and Slices
#⌨️ (2:17:20) Maps and Structs
#⌨️ (2:48:00) If and Switch Statements
#⌨️ (3:21:17) Looping
#⌨️ (3:41:34) Defer, Panic, and Recover
#⌨️ (4:03:57) Pointers
#⌨️ (4:21:30) Functions
#⌨️ (4:57:59) Interfaces
#⌨️ (5:33:57) Goroutines
#⌨️ (6:05:10) Channels
