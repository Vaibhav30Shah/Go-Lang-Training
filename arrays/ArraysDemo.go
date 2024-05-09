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
	fmt.Printf("arr5: %#v\n", arr5)

	//array copy
	arr6 := [3]int{1, 2, 3}
	//arr7 := arr6
	var arr7 = [3]int{}
	arr7 = arr6 //no copy created
	//arr7:=arr6 //no copy is created
	fmt.Printf("arr6==arr7? %v\n", arr6 == arr7) //true
	arr6[0] = -1                                 //changing in arr6 will not change in arr7
	fmt.Printf("arr6==arr7? %v\n", arr6 == arr7) //false

	//array keyed elements
	//we can specify the elements at particular indexes.
	arr8 := [3]int{
		2: 0,
		1: 2,
		0: 8, //here the comma is mandatory otherwise compiler error
	}
	fmt.Printf("arr8: %v\n", arr8)

	//we can also specify only particular index value. For eg only for index 2, rest will be 0(in case of int or float), nil in case of String and false in bool
	arr9 := [3]string{2: "hi"}
	fmt.Printf("arr9: %v\n", arr9)

	//if we specify some indexes and not all, the unspecified indexes will be added at the last(after the index of its current predescessor in case of ...)
	arr10 := [...]int{
		0: 10,
		2, //at arr10[1]
		//2, //error becoz we have no index between 1 and 2
		2: 100,
	}
	fmt.Printf("arr10: %#v\n", arr10)
}
