package main

import "testing"

func TestSelectSort(t *testing.T) {
	values := []int{9, 1, 2, 8, 3, 7, 4, 6, 5}
	result := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	values = selectSort(values)
	for i := 0; i < len(values); i++ {
		if values[i] != result[i] {
			t.Errorf("incorrect sort, %v != %v", values[i], result[i])
		}
	}
}

func TestBubbleSort(t *testing.T) {
	values := []int{9, 1, 2, 8, 3, 7, 4, 6, 5}
	result := []int{1, 2, 3, 4, 5, 6, 7, 7, 9}
	values = bubbleSort(values)
	for i := 0; i < len(values); i++ {
		if values[i] != result[i] {
			t.Errorf("incorrect sort, %v != %v", values[i], result[i])
		}
	}
}
