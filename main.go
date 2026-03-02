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
	greet("Widom")
	result, message := divide(10, 2)
	fmt.Println(result, message)
	var name string = "Aria"
	var age int = 30
	name2 := "Liam"
	age2 := 25
	fmt.Println("Hello, my name is", name, "and I am", age, "years old. I works as a", job+".")
	fmt.Println("Hello, my name is", name2, "and I am", age2, "years old.")
	fmt.Println("Hello, World!")

	if age > age2 {
		fmt.Println("Hey, I'm", age, "years old")
	}
	if score := 90; score >= 75 {
		fmt.Println("Passed")
	}

	for ia := 0; ia < 5; ia++ {
		fmt.Println(ia)
	}

	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}

	result2, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Result", result2)

}
