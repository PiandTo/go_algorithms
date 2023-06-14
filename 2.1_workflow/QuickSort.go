package main

import (
	"fmt"
	"time"
	"math/rand"
)

func partition(arr []int) int {
	low := 0
	high := len(arr) - 1
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if (arr[j] <= pivot) {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		return 
	}
	i := partition(arr)
	quicksort(arr[0:i])
	quicksort(arr[i+1:len(arr)])
}

func main() {
	rand.Seed(time.Now().UnixNano())

    // Get the number of items and maximum item value.
    var num_items, max int;
    fmt.Printf("# Items: ")
    fmt.Scanln(&num_items)
    fmt.Printf("Max: ")
    fmt.Scanln(&max)

    // Make and display the unsorted array.
    arr := make_random_array(num_items, max)
    print_array(arr, 40)
    fmt.Println()

    // Sort and display the result.
    quicksort(arr)
    print_array(arr, 40)

    // Verify that it's sorted.
    check_sorted(arr)
}

func check_sorted (arr []int) {
	for i := 1; i < len(arr); i++ {
		if (arr[i - 1] > arr[i]) {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}
	fmt.Println("The array is sorted")
}

func make_random_array(num_items, max int) []int {
	arr := make([]int, num_items)

	for i := 0; i < num_items; i++ {
		arr[i] = rand.Intn(max)
	}
	return arr
}

func print_array(arr []int, num_items int) {
	if num_items > len(arr) {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:num_items])
	}
}