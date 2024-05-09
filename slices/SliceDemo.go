package slices

import "fmt"

func SliceDemo() {
	//default size of a slice is nil. It is instantiated at rutime
	var slc1 []int
	fmt.Printf("Elements: %#v and size: %v\n", slc1, len(slc1))

	var s []int
	println(s == nil) //true because it is declared slice. and declared slice is nil

	s2 := []int{}
	println(s2 == nil) //false because it is initialized slice and initialized slice has {} AND not nil
	fmt.Printf("Elements: %#v and size: %v\n", s2, len(s2))

	//slice declarations
	slc2 := []int{1, 2}
	fmt.Printf("Elements: %#v and size: %v\n", slc2, len(slc2))

	slc3 := make([]int, 2)
	fmt.Printf("Elements: %#v and size: %v\n", slc3, len(slc3))

	type typeslc []string
	slc4 := typeslc{"A", "B"}
	fmt.Printf("Elements: %#v and size: %v\n", slc4, len(slc4))

	//iterating over a slice
	for index, value := range slc4 {
		fmt.Printf("Index: %d, value: %v\n", index, value)
	}

	/*
			* Comparing Slices
			* Slices can be compared only by nil. So otherwise the slices cant be compared.
			* To compare 2 slices use a for loop to iterate over the slices and compare element by element
		    * It's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)
	*/
	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	var eq bool = true
	if len(a) != len(b) {
		eq = false
	}

	for i, valueA := range a {
		if valueA != b[i] {
			eq = false // don't check further, break!
			break
		}
	}

	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}

	//appending to a slice.
	//Append does not modify the original slice but creates a new copy
	new_app := append(slc4, "hi", "Hello")
	fmt.Printf("Appended slice: %v\nOriginal slice: %v\n", new_app, slc4)

	//appending the elemets of any 1 slice to other
	temp_slc := []int{1, 2, 3}
	slc3 = append(slc3, temp_slc...)
	fmt.Printf("Appended slice: %v\nOriginal slice: %v\n", temp_slc, slc3)

	//copying 1 slice into other
	src := []int{1, 2, 3}
	dst := make([]int, 2) //2nd argument is the length of the slice which is to be made. If set to 0, nothing will be copied
	nn := copy(dst, src)  //returns the length of smaller slice or the elements that are copies from src. to dest.
	fmt.Printf("Source slice: %v\nDesitination slice: %v\nFinal Slice: %v\n", src, dst, nn)

	//slice expression
	//slice_name[start:stop]
	slc5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slc5[1:4]) //2 3 4 -- This is also applicable in terms of array
	fmt.Println(slc5[:3])  //prints from 0 to index=3(3 not included) i.e, 0,1,2 indexes
	fmt.Println(slc5[2:])  //prints from index 2(2 inclusive) to len(slc5)
	fmt.Println(slc5[:])   //prints from index 0 to len(slc5)

	slc5 = append(slc5[:3], 1000) //by doing this the elements after index 3 are deleted and 1000 will be added at index 3
	fmt.Println(slc5)

	slc5 = append(slc5[:3], 2000) //1000 is replaced by 2000
	fmt.Println(slc5)

}
