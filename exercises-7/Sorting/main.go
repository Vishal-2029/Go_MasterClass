package main

import (
	"fmt"

)

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				
			}

			
		}
	}
	return arr
}

 func selectsort(Arr []int )[] int {
	for i := 0; i <len(Arr); i++ {
		smallidx := i
		for j := i+1 ; j < len(Arr) ; j++ {
			if Arr [j] < Arr [smallidx]{
				smallidx = j
			}

		}
		Arr[i], Arr[smallidx] = Arr[smallidx], Arr[i]
	}
	return Arr
}

func insertionsort(ARR []int) []int{
		for i := 1; i < len(ARR); i ++ {
			curr := ARR[i]
			prev := i-1

			for prev >= 0 && ARR[prev] > curr {
				ARR[prev+1] = ARR[prev]
				prev--
			}

			ARR[prev+1] = curr
		}
		return ARR
}


func main() {
	arr := []int{9, 5, 4, 7, 2, 6, 4, 5, 3, 1}
	Arr := []int{9, 5, 4, 7, 2, 4, 5, 9, 1, 3}
	ARR := []int{9, 6, 4, 5, 3, 4, 5, 9, 1, 3}
	sortedArr := bubbleSort(arr)
	selectsortArr := selectsort(Arr)
	insertionsortArr := insertionsort(ARR)

	fmt.Print("Bubble Sorted array: ")
	for _, v := range sortedArr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Select Sorted array: ")
	for _, v := range selectsortArr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Insertion Sorted array: ")
	for _, v := range insertionsortArr {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
