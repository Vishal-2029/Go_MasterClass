package main

import "fmt"



func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main(){
	var n []int
	var N int

	fmt.Print("Enter the Number: ")
	fmt.Scan(&N)

	for i := 0; i <= N; i++ { 
		if isPrime(i) {
			n = append(n, i)
		}
	}
	fmt.Println(n)
}