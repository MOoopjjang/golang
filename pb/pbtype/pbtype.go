package pbtype

import "fmt"

type User struct {
	Name    string
	Number  string
	Address string
}

func (u *User) ToString() string {
	return fmt.Sprintf("name = %s , Number = %s , Address = %s", u.Name, u.Number, u.Address)
}

func (u *User) Edit(cNumber string, cAddress string) {
	u.Number = cNumber
	u.Address = cAddress
}

type PhoneBook struct {
	Count int
	Data  []User
}
