package methods

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.Name, u.Age)
}

func (u *User) HaveBirthday() {
	u.Age++
}

