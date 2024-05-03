package flow_control

import "fmt"

func ForLoopDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	println()

	//for (i := 0;i < 10;i++) { //() gives compiler error
	//	fmt.Println(i)
	//}

	//for i := 0; i < 10; { //this becomes infinite loop
	//	fmt.Println(i)
	//}

	for i := 0; i < 10; {
		fmt.Println(i)
		i++ //we can write the inc/dec part inside
	}

	println()

	//while loop from for loop
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	//for i := 0; ; i++ { //infinite loop
	//	fmt.Println(i)
	//}

	println()
name:
	y := 10
	_ = y
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //skips even numbers
		} else {
			for j := 0; j < 10; j++ {
				if j == i {
					print(j)
					break //breaks inner loop
				}
				//println(i, j)
			}
			break //breaks outer loop
		}
		goto name
	}

}
