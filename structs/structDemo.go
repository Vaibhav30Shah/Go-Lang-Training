package structs

import "fmt"

func StructDemo() {
	//creating structs
	type book struct {
		title  string
		author string
		price  float64
	}

	//fmt.Println(book) //error. we cant print the struct definition
	myBook := book{"The hound of the baskervilles'", "Sir arthur conan doyle", 559.65} //not recommended way of creating structs
	fmt.Println(myBook)
	fmt.Printf("%T\n", myBook) //structs.book type

	myBook2 := book{title: "Sherlock Homes", price: 499.79, author: "Sir Arthur Conan Doyle"}
	fmt.Println(myBook2)

	//retrieving values from struct
	//use . Dot(.) is a selection operator for structs
	fmt.Printf("%v\n", myBook2.title)
	//we can also use . to update the value of a struct

	//comparison of structs
	fmt.Println(myBook2 == myBook2) //true because fields are same only.

	newBook := myBook2
	newBook.title = "njbhjbch"

	fmt.Println(newBook, myBook) //updating the new struct will not update the original list

	//Anonymous structs and anonymous struct fields
	anon := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "John",
		lastName:  "Smith",
		age:       25,
	}
	fmt.Println(anon)
}
