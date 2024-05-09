package slices

import "fmt"

func BackingArray() {
	//when any one slice is changed, its all corresponding slices are changed
	// as the slice elements are stored in backing array and hence the backing array is only 1 for similar slices(same name)
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s1, s2 := slc[1:4], slc[2:6]
	fmt.Printf("Before %v: \n", slc)

	s1[2] = 2000
	s2[3] = 3000
	fmt.Printf("After %v: \n", slc)
	fmt.Printf("s1 %v s2 %v\n", s1, s2)

	//backing array demo
	backArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s3, s4 := backArr[1:4], backArr[2:7]
	fmt.Printf("Before(Array) %v: \n", backArr)

	s3[2] = 9000
	s4[3] = 8888

	fmt.Printf("After(Array) %v: \n", backArr) //backArr changed
	fmt.Printf("s3 %v s4 %v\n", s3, s4)

	//Slices are not modified when used by append
	s5 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s6 := []int{}
	s6 = append(s6, s5[2:6]...) //... is mandatory
	s6[0] = 8888
	fmt.Printf("s5 %v and s6 %v\n", s5, s6) //s5 not modified

	s7 := append(s5[:4], 1000, 2000)
	fmt.Printf("s7 %v len %v cap %v\n", s7, len(s7), cap(s7)) //cap is original size of backing array
	//fmt.Printf("s7[6]: %v\n", s7[6])                          //panic: index out of range
	fmt.Printf("s7[6]: %v\n", s7[6:7]) //gives result as slicing works on backing array
}
