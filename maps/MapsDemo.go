package maps

import "fmt"

func MapsDemo() {
	//var m1 map[string]string
	//m1["one"] = "one"
	//fmt.Println(m1) //panic: assignment to entry in nil map

	//use make() function for complex ds like maps
	m1 := make(map[string]string) //make function allocates and initializes a hashmap ds and returns a map value that points to it
	fmt.Println(m1)               //empty map
	m1["one"] = "one"
	fmt.Println(m1) //no panic

	//m2 := make(map[[4]int]int) //slice as a key is not allowed as key must be a comparable type

	//to delete an entry a delete function is used
	delete(m1, "one")
	fmt.Println(m1)

	//comparing maps
	//a map can only be compared to nil.
	//we can convert the maps to string and then compare using Sprintf()
	m2, m3 := map[int]int{1: 1, 2: 2}, map[int]int{1: 1, 2: 2}
	//fmt.Println(m2==m3) //maps can only be compared to nil
	s2 := fmt.Sprintf("%s", m2)
	s3 := fmt.Sprintf("%s", m3)
	println(s2 == s3) //true

	//Internal representation of maps
	//Map header and cloning maps
	//When declaring a map, go creates a pointer to the map header in memory. In memory only map header is stored and not the whole map.
	//When we copy a map, the whole internal ds isn't copied. Just referenced

	m4 := map[string]int{"A": 10, "B": 20}
	m5 := m4
	m5["A"] = 9999
	fmt.Println(m4) //m4 also modified

	//cloning a map
	m6 := make(map[string]int)
	for key, value := range m4 {
		m6[key] = value
	}
	m6["A"] = 98986321
	fmt.Println(m6, m4) //m4 not modified

}
