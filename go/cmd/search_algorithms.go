package main

import "fmt"

func Linear_search_example() {
	arr := []int{1, 45, 99, 23, 5, 67, 89}
	// O(n) - Linear Time
	linear_search(arr, 23)
	// O(1) - Constant Time
	linear_search(arr, 1)
	// O(n) - Linear Time Worst Case.
	linear_search(arr, 89)
}

func linear_search(arr []int, target int) bool {
	for i := range arr {
		if arr[i] == target {
			return true
		}
	}
	return false
}

func Binary_search_example() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	// O(log n) - Logarithmic Time
	binary_search(arr, 70)
}
func binary_search(arr []int, target int) bool {
	iter := 0
	left := 0
	right := len(arr) - 1
	for left <= right {
		iter++
		mid := (left + right) / 2
		fmt.Printf("\nIteração %d: left=%d, right=%d, mid=%d, arr[mid]=%d\n", iter, left, right, mid, arr[mid])
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func Challenge_binary_search() bool {
	arr := []string{"Arroz", "Batata", "Carne"}
	iter := 0
	// vou buscar arroz
	target := "Arroz"
	left := 0
	right := len(arr) - 1
	for left <= right {
		iter++
		mid := (left + right) / 2
		fmt.Printf("\nIteração %d: left=%d, right=%d, mid=%d, arr[mid]=%s\n", iter, left, right, mid, arr[mid])
		if arr[mid] == target {
			return true
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
