package main 

func Bubblesort(arr []int) []int {
	for i:= 0; i < len(arr) ; i++ {
		for j:= 0; j < len(arr)-i-1; j++ {
			if arr [j] > arr [j+1] {
				arr [j], arr [j+1] = arr [j+1], arr[j]
			}
		}
	}
	return arr
}

func main(){
	arr := []int{9,5,4,7,2,6,4,5,3,4,5,9,1,3}

	sortedArr := Bubblesort(arr)

	print("Sorted array: ")
	for _,v := range sortedArr{
		print(v, " ")
	}
	println()
}