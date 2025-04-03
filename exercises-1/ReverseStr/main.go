package main

import "fmt"

func main() {
	str := "Vishal"

	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)
}

func reverse(s string) string {
	reversed := ""
	for _, char := range s {
		reversed = string(char) + reversed
	}
	return reversed
}