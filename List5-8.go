package main

import "fmt"

type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "さようなら"
	p.name = "Bob"
}

func main() {
	p := person{}
	i := 2
	s := "こんにちは"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}
