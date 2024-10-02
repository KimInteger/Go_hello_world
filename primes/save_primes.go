// file_utils.go
package main

import (
	"os"
	"strconv"
	"strings"
)

func SavePrimesToFile(filename string, primes []int) error {
	var primesStr []string

	for _, prime := range primes {
		primesStr = append(primesStr, strconv.Itoa(prime)) // 정수를 문자열로 변환
	}
	primesOutput := strings.Join(primesStr, ", ") // 문자열로 조인

	err := os.WriteFile(filename, []byte(primesOutput), 0644) // 0644는 파일 권한

	if err != nil {
		return err
	}
	return nil
}
