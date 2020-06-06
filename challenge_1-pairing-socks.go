// John works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.

// For example, there are  socks with colors . There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .

// Function Description

// Complete the sockMerchant function in the editor below. It must return an integer representing the number of matching pairs of socks that are available.

// sockMerchant has the following parameter(s):

// n: the number of socks in the pile
// ar: the colors of each sock
// Input Format

// The first line contains an integer , the number of socks represented in .
// The second line contains  space-separated integers describing the colors  of the socks in the pile.

// Constraints

//  where 
// Output Format

// Return the total number of matching pairs of socks that John can sell.

// Sample Input
// 9
// 10 20 20 10 10 30 50 10 20

// Sample Output
// 3

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
