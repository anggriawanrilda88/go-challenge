package main

import (
	"fmt"
)

func contains(arr []int32, value int32) bool {
	for _, a := range arr {
		if a == value {
			return true
		}
	}
	return false
}

func RemoveIndex(s []int32, num int32) []int32 {
	for i, val := range s {
		if val == num {
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var PairingSock []int32
	allPairingSock := int32(0)
	for i, num := range ar {
		if i == 0 {
			PairingSock = append(PairingSock, num)
		} else {
			if contains(PairingSock, num) {
				PairingSock = RemoveIndex(PairingSock, num)
				allPairingSock = allPairingSock + 1
			} else {
				PairingSock = append(PairingSock, num)
			}
		}
	}
	return allPairingSock
}

func main() {
	n := int32(9)
	ar := []int32{6, 5, 2, 3, 5, 2, 2, 1, 1, 5, 1, 3, 3, 3, 5}
	fmt.Println(sockMerchant(n, ar))
}
