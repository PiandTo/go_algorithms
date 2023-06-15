package main

import (
	"fmt"
	"time"
	"math/rand"
	"strconv"
)

type Customer struct {
	id	string
	num_purchases	int
}

func make_random_array(num_items, max int) []Customer {
	arr := make([]Customer, num_items)
	for i := 0; i < num_items; i++ {
		arr[i].id = strconv.Itoa(i) + "C"
		arr[i].num_purchases = rand.Intn(max);
	}
	return arr
}

func print_array(arr []Customer, num_items int) {
	if num_items > len(arr) {
		fmt.Println(arr)
	} else {
		fmt.Println(arr[:num_items])
	}
}

func check_sorted(arr []Customer) {
	for i := 1; i < len(arr); i++ {
		if (arr[i - 1].num_purchases > arr[i].num_purchases) {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}
	fmt.Println("The array is sorted")
}

func counting_sort(cust []Customer, max int) []Customer {
	arr := make([]int, max)
	result := make([]Customer, len(cust))
	for i := 0; i < len(cust); i++ {
		arr[cust[i].num_purchases] = arr[cust[i].num_purchases] + 1
	}
	fmt.Println(arr)
	for j := 1; j < len(arr); j++ {
		arr[j] = arr[j] + arr[j - 1]
	}
	fmt.Println(arr)
	for k := max - 1; k > -1; k-- {
		arr[cust[k].num_purchases] = arr[cust[k].num_purchases] - 1
		result[arr[cust[k].num_purchases]] = cust[k]
	}
	return result
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
    sorted := counting_sort(arr, max)
    print_array(sorted, 40)

    // Verify that it's sorted.
    check_sorted(sorted)
}