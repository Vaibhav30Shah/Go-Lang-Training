package methods

import "fmt"

// declaring a new type
type cars []string

func (n cars) printCars() {
	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.print()

	for index, value := range n {
		fmt.Println(index, value)
	}
}

type changes struct {
	brand string
	price int
}

func (c changes) printChangesByValue(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// method with a pointer receiver
func (c *changes) printChangesByPointer(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func MethodDemo() {
	suvs := cars{"Tata Punch", "Creta", "M Thar"}
	suvs.printCars()

	//diff bw method and function is that a method is for any particular type. while a function is for package

	first := changes{"Tata", 100}
	fmt.Println(first.brand, first.price)

	first.printChangesByValue("Hyundai", 10)
	fmt.Println(first.brand, first.price) // no change

	first.printChangesByPointer("Toyota", 20)
	fmt.Println(first.brand, first.price) //changed to toyota 20

	// calling the method with a pointer receiver. It changes the value!
	(&first).printChangesByPointer("Seat", 25000)
	fmt.Println(first) // -> {Seat 25000}

	// declaring a pointer to car
	var yourCar *changes
	yourCar = &first      // it stores the address of myCar
	fmt.Println(*yourCar) // -> {Seat 25000}

	// calling the method on pointer type
	// valid ways to call the method:
	yourCar.printChangesByPointer("VW", 30000)
	(*yourCar).printChangesByPointer("VW", 30000)
	fmt.Println(*yourCar) // => {VW 30000}

	// in the above example both myCar and yourCar variables have been modified.
	fmt.Println(first) // => {VW 30000}
}
