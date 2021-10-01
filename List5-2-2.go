package main

import (
	"errors"
	"fmt"
	"os"
)

func divAndRemainder(numerator, denominator int) (result int, remainder int, err error) {
	// 適当に値を代入
	result, remainder = 20, 30
	if denominator == 0 {
		return 0, 0, errors.New("0での除算はできません")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
