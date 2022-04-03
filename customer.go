package main

import "fmt"

type customer struct {
	FName    string
	LName    string
	Username string
	Password string
	Email    string
	Phone    int
	Address  string
}

func userCredentials(c customer) (string, string) {
	fmt.Println(c.Username, c.Password)
	return c.Username, c.Password
}

func userAddress(c customer) string {
	fmt.Println(c.Address)
	return c.Address
}

func allInfo(c customer) {
	fmt.Println(c.FName, c.LName, c.Username, c.Password, c.Email, c.Phone, c.Address)
	return c.FName, c.LName, c.Username, c.Password, c.Email, c.Phone, c.Address
}

func main() {

	var customer_1 = customer{FName: "Micheal", LName: "Jordan", Username: "MJ2020", Password: "1234567", Email: "MJ2020@gmail.com", Phone: 12345678, Address: "18227 Capstan Greens Road Cornelius, NC 28031"}

	fmt.Println(customer_1.allInfo())
	fmt.Println(customer_1.userCredentials())
	fmt.Println(customer_1.userAddress())
}

//Declare a struct

//Create Methods
