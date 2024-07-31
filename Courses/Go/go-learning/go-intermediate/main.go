package main

import "fmt"

// var name struct {
// 	value int
// 	num float32
// 	str string
// }

// type person struct {
// 	age int 
// 	name string
// }

// func (p *person) changeUser() {
// 	p.age = 20
// 	p.name = "Jude"
// }

// // func showUser(p *person) {
// // 	p.age = 22
// // 	p.name = "Maxwill"
// // 	fmt.Println(*p)
// // }


// func main() {
// 	// name.value = 2
// 	// name.num = 3.14
// 	// name.str = "name"
// 	// fmt.Println(name)
// 	var user person = person{20, "Meezaan"}
// 	user.changeUser()
// 	fmt.Println(user)
// 	// showUser(&user)
// }

// -----------------------------------------------------------------------------

// type Car string
// type Motorcycle string

// type Drive interface {
// 	Steer()
// }

// func (c Car) Steer() {
// 	fmt.Println("Car can steer")
// }

// func (m Motorcycle) Steer() {
// 	fmt.Println("Motorcycle can steer")
// }

// func canSteer(d Drive) {
// 	d.Steer()
// }

// func main() {
// 	canSteer(Car("Car"))
// 	canSteer(Motorcycle("Motorcycle"))
// }

// -----------------------------------------------------------------------------

func doSomething() {
	defer fmt.Println("Deferred")
	fmt.Println("1")
	panic(doSomething)
	fmt.Println("2")
}

func main() {
	doSomething()
	fmt.Println("After doSomething")
}
