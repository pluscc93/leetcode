package main

import (
	"fmt"
)

func compress(arr []string) []string {
	var index int
	for i := 0; i < len(arr); i++ {
		count := 1
		for i+1 < len(arr) && arr[i] == arr[i+1] {
			fmt.Println(i)
			count++
			i++
		}
		arr[index] = arr[i]
		index++
		if count > 1 {
			arr[index] = fmt.Sprintf("%d", count)
			index++
		}
	}
	return arr[:index]
}
func main() {
	array := []string{"a", "a", "a", "a", "b", "b", "c", "c", "c", "c"}

	fmt.Println(compress(array))
	fmt.Println(array)

}
