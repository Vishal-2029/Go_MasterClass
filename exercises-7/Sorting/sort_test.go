package main

import (
	"math/rand"
	"testing"
	"time"
)


func generateRandomSlice(n int) []int {
	rand.Seed((time.Now().UnixNano()))
	slice := make([]int, n)
	for i := range slice {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := generateRandomSlice(100)
		bubbleSort(arr)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Arr := generateRandomSlice(100)
		selectsort(Arr)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ARR := generateRandomSlice(100)
		insertionsort(ARR)
	}
}
