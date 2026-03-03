package main

import "fmt"

func basicDataTypes() {
	var name string = "Aria"
	var age int = 30
	name2 := "Liam"
	age2 := 25
	fmt.Println("Hello, my name is", name, "and I am", age, "years old. I works as a", job+".")
	fmt.Println("Hello, my name is", name2, "and I am", age2, "years old.")

	if age > age2 {
		fmt.Println("Hey, I'm", age, "years old")
	}
}

func basicIfAndLoops() {
	if score := 90; score >= 75 {
		fmt.Println("Passed")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	count := 0
	for count < 5 {
		fmt.Println(count)
		count++
	}
}
