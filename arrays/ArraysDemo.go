package arrays

import "fmt"

func ArraysDemo() {
	var arr [4]string

	//println(arr) //error
	fmt.Printf("arr = %v\n", arr)
	fmt.Printf("len(arr) = %v\n", len(arr))

	//array declaration types
	var arr1 [4]int
	fmt.Printf("arr1 = %v\n", arr1)

	type nn = int8
	var n nn = 5
	arr2 := [4]int8{1, 2, 3, n}
	fmt.Printf("arr2 = %v\n", arr2)

	var arr3 = [4]float64{}
	fmt.Printf("arr3 = %v\n", arr3)

	var arr4 = [...]string{"hi", "hello", "world"} //... is ellipses operator and we can add n no of items with this
	fmt.Printf("arr4 = %v\n", arr4)
	fmt.Printf("len(arr4) = %v\n", len(arr4))

	//operations in array
	//for each type loop
	for index, value := range arr2 {
		fmt.Printf("arr2[%d] = %v\n", index, value)
	}

	//multi dimensional array
	arr5 := [3][][2]int{}
	fmt.Printf("arr5: %#v", arr5)
}
