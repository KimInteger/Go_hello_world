package main

import (
	"fmt"
	"primes/find"
)

func main() {
	var start, end int

	fmt.Print("소수를 찾고자 하는 시작 숫자를 입력하세요! 2이상! 2이하로 적으면 2가됩니다!")
	fmt.Scan(&start)
	fmt.Print("어디까지 찾을래요?")
	fmt.Scan(&end)

	primesList := find.FindPrimesInRange(start, end)

	fmt.Printf("%d부터 %d까지의 소수는 %v\n입니다", start, end, primesList)

	// 결과를 파일에 저장
	err := SavePrimesToFile("primes.txt", primesList)
	if err != nil {
		fmt.Println("파일 저장 오류:", err)
		return
	}
	fmt.Println("소수 목록이 파일에 저장되었습니다.")
}
