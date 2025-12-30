package main

import "fmt"

func main() {

	a := 10
	b := 20
	fmt.Println("This sum of two number", a*b)

	Sample()
	go GoRoutine()
}

func GoRoutine() {
	fmt.Println("I am goroutine...!!!")

}

func Sample() {
	fmt.Println("I am go function")
}
