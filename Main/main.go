package main

import (
	"fmt"
	_ "fmt"
	"time"
)

func main() {
	fmt.Printf("Program Started\n")
	//basics.F()
	//variables.VariableDeclaration()
	//variables.TypesOfVariables()
	//datatypes.ArraysAndSlices()
	//datatypes.Maps()
	//flow_control.GettingUserInput()
	//flow_control.IfElseIfDemo()
	//flow_control.ForLoopDemo()
	//flow_control.SwitchDemo()
	//arrays.ArraysDemo()
	//slices.SliceDemo()
	//slices.BackingArray()
	//strings.StringsDemo()
	//maps.MapsDemo()
	//structs.StructDemo()
	//functions.FunctionDemo()
	//functions.DeferStatement()
	//pointers.PointerDemo()
	//methods.MethodDemo()
	//interfaces.InterfaceDemo()
	//goroutines.GoRoutineDemo()
	//concurrency.MainUrlChecker()
	//concurrency.RaceCondition()
	//concurrency.Mutexes()
	//concurrency.ChannelDemo()
	//concurrency.ChannelSimple()
	size := 1000000
	values1 := make([]int32, size)
	st1 := time.Now().UnixMilli()
	for index := 0; index < size; index++ {
		values1[index] = int32(index)
	}
	et1 := time.Now().UnixMilli()
	t1 := et1 - st1
	var values2 []int32
	st2 := time.Now().UnixMilli()
	for index := 0; index < size; index++ {
		values2 = append(values2, int32(index))
	}
	et2 := time.Now().UnixMilli()
	t2 := et2 - st2
	fmt.Println(t1)
	fmt.Println(t2)
}
