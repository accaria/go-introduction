package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func structExercise() {
	user := User{
		Name: "Aria",
		Age:  25,
	}
	fmt.Println("Name is", user.Name)
	fmt.Println("Age is", user.Age)
	updateAge(&user)
	fmt.Println("Age is now", user.Age)
}

func updateAge(u *User) {
	u.Age = 99
}

func createUser(name string, age int) User {
	return User{
		Name: name,
		Age:  age,
	}
}

func structExercise2() {
	u := createUser("Liam", 30)
	fmt.Println(u.Name, u.Age)
}

type Address struct {
	City    string
	Country string
}
type Customer struct {
	Name    string
	Address Address
}

func structExercise3() {
	c := Customer{
		Name: "Aria",
		Address: Address{
			City:    "Kyoto",
			Country: "Japan",
		},
	}
	fmt.Println(c.Address.City)
}

func structExercise4() {
	users := []User{
		{Name: "Aria", Age: 25},
		{Name: "Liam", Age: 30},
	}
	for _, u := range users {
		fmt.Println(u.Name, u.Age)
	}
}
