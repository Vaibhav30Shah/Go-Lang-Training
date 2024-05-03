package flow_control

func IfElseIfDemo() {
	var flag bool = true

	if flag {
		flag = false
	} else {

	}

	if /* ( */ flag /* ) */ { //here () are optional

	} else {

	}

	var age = 18
	_ = age
	//if age { //error
	//
	//}

	if val := 20; val < 18 { //simple if with assignment and comparison both in a single statement
		println(val)
	} else {
		println("Else executed", val)
	}
}
