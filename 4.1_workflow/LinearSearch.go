package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

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

	for {
		// Get the target as a string.
		var target_string string;
		fmt.Printf("Target: ")
		fmt.Scanln(&target_string)

		// If the target string is blank, break out of the loop.
		if len(target_string) == 0 { break }

		// Convert to int and find it.
		target, _ := strconv.Atoi(target_string)

		index, num_tests := linear_search(arr, target)

		if index < 0 {
			fmt.Printf("Target %d not found, %d tests\n", target, num_tests)
		} else {
			fmt.Printf("arr[%d] = %d, %d tests\n", index, arr[index], num_tests)
		}
	}
}

func linear_search(arr []int, target int) (index, num_tests int) {
	for i := 0; i < len(arr); i++ {
		if (arr[i] == target) {
			return i, i + 1
		}
	}
	return -1, len(arr)
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