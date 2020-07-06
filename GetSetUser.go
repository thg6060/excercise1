package main

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func (u *User) GetName() string {
	return u.name
}
func (u *User) GetAge() int64 {
	return u.age
}
func (u *User) GetAddress() string {
	return u.address
}
func (u *User) GetGender() bool {
	return u.gender
}
func MapUser() map[string]*User {
	m := make(map[string]*User)
	us1 := User{name: "Giang1", age: 21, address: "Ha Noi1", gender: true}
	us2 := User{name: "Giang2", age: 21, address: "Ha Noi2", gender: true}
	us3 := User{name: "Giang3", age: 21, address: "Ha Noi3", gender: true}
	us4 := User{name: "Giang4", age: 21, address: "Ha Noi4", gender: true}
	us5 := User{name: "Giang5", age: 21, address: "Ha Noi5", gender: true}
	us6 := User{name: "Giang6", age: 21, address: "Ha Noi6", gender: true}
	us7 := User{name: "Giang7", age: 21, address: "Ha Noi7", gender: true}
	us8 := User{name: "Giang8", age: 21, address: "Ha Noi8", gender: true}
	us9 := User{name: "Giang9", age: 21, address: "Ha Noi9", gender: true}
	us10 := User{name: "Giang10", age: 21, address: "Ha Noi10", gender: true}

	m[us1.GetName()] = &us1
	m[us2.GetName()] = &us2
	m[us3.GetName()] = &us3
	m[us4.GetName()] = &us4
	m[us5.GetName()] = &us5
	m[us6.GetName()] = &us6
	m[us7.GetName()] = &us7
	m[us8.GetName()] = &us8
	m[us9.GetName()] = &us9
	m[us10.GetName()] = &us10

	return m
}
func MaptoSlice(m map[string]*User) []User {
	var sl []User
	for item := range m {
		sl = append(sl, *m[item])
	}
	return sl
}
