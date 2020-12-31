package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) update(n string, a int) {
	p.name = n
	p.age = a
	fmt.Println(p.name, p.age)

}

func display(st string) string {
	return fmt.Sprint("Hello From golang ", st)
}

func main() {
	var a [5]int
	a[0] = 10
	var p Person
	p.update("Ravi", 26)

	mess := display("Welcome to the new world of go programming")
	//fmt.Println(per)
	fmt.Println(mess)
	fmt.Println(a)

	//slice
	//var intSlice []int
	//var strSlice = []string{"India", "Canada", "Japan"}

	//stooges := [...]string{"Moe", "Larry", "Curly"}

	//stooges[3] = "amit"

	init := make([]string, 10)

	init[0] = "ravi"

	for _, v := range init {
		fmt.Printf("%s", v)
	}

}
