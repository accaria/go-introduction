package main

import "fmt"

func pointerPractice() {
	x := 10
	fmt.Println("Value of X is", x)
	p := &x
	fmt.Println("The pointer address for X is", p)
	fmt.Println("The value for X from the address is", *p)
}

func addFive(num *int) {
	*num = *num + 5
}

func pointerExercise() {
	x := 10
	fmt.Println("Before:", x)
	addFive(&x)
	fmt.Println("After", x)
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func pointerExercise2() {
	x := 5
	y := 10
	fmt.Println("Before,", x, y)
	swap(&x, &y)
	fmt.Println("After,", x, y)
}
