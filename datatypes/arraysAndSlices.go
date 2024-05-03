package datatypes

import "fmt"

func ArraysAndSlices() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(arr)
	fmt.Printf("%T", arr)

	var slc = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(slc)
	fmt.Printf("%T", slc)
	ints := append(slc, 2)
	//var arr2 = append(arr, 2)
	fmt.Println(ints)
	fmt.Println(slc)

}
