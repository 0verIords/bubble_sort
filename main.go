package main

import "fmt"

func selectSort(array []int) []int {
	for i := range array {
		temp := i
		j := i + 1
		for j < len(array) {
			if array[j] < array[temp] {
				temp = j
			}
			j++
		}
		temporary := array[i]
		array[i] = array[temp]
		array[temp] = temporary
	}
	return array
}

func bubbleSort(array []int) []int {
	var temporary int
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				temporary = array[j]
				array[j] = array[j+1]
				array[j+1] = temporary
			}
		}
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}

	return array
}

func main() {
	s := []int{2, 4, 8, 1, 0, 12, 3, 9, 5, 7, 6}
	ss := selectSort(s)
	fmt.Println(ss)

	a := []int{2, 4, 8, 1, 0, 12, 3, 9, 5, 7, 6}
	aa := bubbleSort(a)
	fmt.Println(aa)
}
