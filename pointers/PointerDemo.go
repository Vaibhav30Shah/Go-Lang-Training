package pointers

import "fmt"

func PointerDemo() {
	val := 10
	fmt.Printf("pointer value: %v\n", val)
	fmt.Printf("pointer address: %v\n", &val) //& -> addressOf operator

	// pointer declaration
	//1.
	var x int = 20
	ptr := &x                                                     //it is a pointer to int i.e, *int. Pointer holds address of a varible
	fmt.Printf("pts Type: %T, Value: %v\n", ptr, ptr)             //value of a pointer is an address of varible(here x)
	fmt.Printf("Address of pointer: %p and of x: %v\n", &ptr, &x) //value of pointer = address of x. and address of pointer is different
	fmt.Printf("Value accessing through a pointer: %v\n", *ptr)   //ptr=&x => *ptr=*&x => *ptr=x

	//2.
	var ptr2 *float64                                    //uninitialized pointer. zero value of a pointer is nil
	fmt.Printf("ptr2 Type: %T, Value: %v\n", ptr2, ptr2) //value is nil

	//3.
	ptr3 := new(string)
	str := "Hi"
	ptr3 = &str
	fmt.Printf("ptr3 Type: %T, Value: %v\n", ptr3, ptr3)

	//dereferencing of pointers(* -> dereferencing operator)
	*ptr3 = "Hello"
	fmt.Printf("ptr3: %v, str: %v\n", (*ptr3), str) //str also changes to hello as it is dereferenced above

	//To declare pointer to a pointer we can do following
	a := 10
	ptr4 := &a
	ptrptr4 := &ptr4
	fmt.Printf("ptr4: %v, ptrptr4: %v\n", ptr4, ptrptr4)
	fmt.Printf("Value of *ptrptr4: %v and of **ptrptr4: %v\n", *ptrptr4, **ptrptr4) //*ptrptr4 contains the address of a i.e, value of ptr4

	//comparing pointers
	//2 pointers are equal if both point to the same variable or are nil
	aa := 5
	ptr5 := &aa

	bb := 5
	ptr6 := &bb

	fmt.Println(ptr5 == ptr6) //false because both points to diff memory locations

	ptr7 := &bb
	fmt.Println(ptr6 == ptr7) //true because both point to same variable i.e., memory location

	//passing and returning a pointer in a function
	cc := 10
	ptr8 := &cc

	fmt.Printf("Before: %v, Location: %v\n", cc, &cc) //10
	change(ptr8)
	fmt.Printf("After: %v, Location: %v\n", cc, &cc) //9999, location remains same of both
}

func change(a *int) *int {
	*a = 9999
	return a
}
