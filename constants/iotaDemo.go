package constants

import "fmt"

func IotaDemo() {
	const (
		v1 = iota
		v2 = iota
		v3 = iota
	)
	fmt.Println(v1, v2, v3) //0 1 2

}
