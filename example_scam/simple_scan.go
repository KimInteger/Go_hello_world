package main

import (
	"fmt"
)

func main() {
	var b string

	fmt.Print("입력해주세요: ")
	_, err := fmt.Scan(&b)
	if err != nil {
		fmt.Println("입력 오류:", err)
		return
	}

	fmt.Println(b)
}
