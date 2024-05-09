package strings

import (
	"fmt"
	"strings"
)

func StringsDemo() {
	str1 := "Hello"
	fmt.Println(str1)
	fmt.Printf(`"Hello!" How Are you?`)    //raw String
	fmt.Printf(`"Hello!" \n How Are you?`) //no effect of \n inside `.

	fmt.Println()
	//multiline strings? use `
	fmt.Printf(`"Hello!"
How Are you?`)
	fmt.Println()
	//it is the same as
	fmt.Println("\"Hello\"! \nHow are you?")

	//strings, characters, bytes and runes
	char1, char2 := 'a', 'A'
	fmt.Printf("Type: %T, value: %v\n", char1, char2) //int32 and 65

	str2 := "Motadata"
	fmt.Printf("Byte at pos 1: %v and its type: %T\n", str2[0], str2[0]) //byte(uint8) is displayed and not rune(int32)
	fmt.Printf("Length of String: %v, %T\n", len(str2), len(str2))

	for i, value := range str2 {
		fmt.Printf("Index: %v, Value: %c\n", i, value) //%c for displaying character/rune value
	}

	//string slicing
	//the sliced and the original strings shares the same backing array
	//slicing returns the bytes and not runes
	str3 := "Lok Sabha General Elections 2024"
	fmt.Printf("Sliced String: %v, Type: %T\n", str3[2:11], str3[2:11])

	//how to get the runes and not bytes via slicing?
	str4 := "64 देशों में चुनाव, सबको भारत से सीखना चाहिए... अहमदाबाद में वोट डालकर बोले PM मोदी"
	fmt.Printf("Sliced String: %v, Type: %T\n", str4[2:11], str4[2:11])

	s := []string{"Hi", "Bye"}
	strs := strings.Join(s, " ")
	println(strs)

	sttt := "..Hi! Hello!!"
	fmt.Println(strings.Trim(sttt, "!"))
}
