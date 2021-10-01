package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string	// 名
		LastName  string	// 姓
		Age       int		// 年齢
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)

	// 姓（LastName）でソート
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)

	// 年齢（Age）でソート
	sort.Slice(people, func(i int, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println(people)
}
