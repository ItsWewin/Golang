package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)
	foo(array[:])
	fmt.Println(array)

	lisa := &user{"Lisa", "Lisa@email.com", 12, false}
	lisa.notify()
	fmt.Println((*lisa).name)

	lisa.changeEmail("LisaNew@email.com")
	fmt.Println((*lisa).email)
}

func foo(slice []int) {
	slice[1] = 100
}

func (u user) notify() {
	u.name = "lisa2"
	fmt.Printf("Sending User Email to %s, %s\n",
		u.name,
		u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}
