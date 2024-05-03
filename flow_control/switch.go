package flow_control

func SwitchDemo() {
	choice := 3

	switch choice { //here also () are optional; here condition is even
	case 1: //here also even must be there
		println("Hi")
		break //break is optional go adds break on its own

	case 2, 3: //it acts as a OR
		println("Bye")

	default:
		println("Invalid choice")
	}

	n := 2
	switch true { //here bool. missing expresion like switch{} means switch true{}
	case n%2 == 0: //here also bool

		println("Even")
	case n%2 != 0:
		println("Odd")
	}
}
