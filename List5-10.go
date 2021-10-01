package main

import "fmt"

func modMap(m map[int]string) {
	m[2] = "こんにちは"
	m[3] = "さようなら"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func main() {
	m := map[int]string{
		1: "一番目",
		2: "二番目",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}
