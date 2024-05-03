package datatypes

import (
	"fmt"
	"strconv"
)

func Maps() {
	var maps map[int]string
	maps = make(map[int]string)
	maps[1] = "Hello"
	maps[2] = "World"
	maps[3] = "Horse"
	maps[4] = "Gopher"
	fmt.Println(maps)

	var a uint32
	var b uint
	_, _ = a, b
	//fmt.Println(a==b)

	var str = fmt.Sprintf("%f\n", 44.44)
	fmt.Printf("%T\n", str)

	sc, err := strconv.ParseFloat("Ab", 32)
	fmt.Println(sc, err)

	var abc rune
	var def int32
	println(abc == def)
}
