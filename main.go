package main

import "fmt"

var job string = "Mechanic"

func greet(name string) {
	fmt.Println("Hello", name)
}

func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "Cannot divide by zero"
	}
	return a / b, "Success"
}

func safeDivide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	structExercise()
	structExercise2()
	structExercise3()
	structExercise4()
}
