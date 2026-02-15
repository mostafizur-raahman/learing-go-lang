package main

import "fmt"

var a int
var b int = 10
var c, d int

// m := 100 // this is only inside function
// f, g := 10, "Hello"

var (
	j string
	k float64
)

func main() {
	var ok string
	fmt.Println(ok)

	arr := [10]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
	}

	subArry := arr[:3]
	fmt.Println(subArry)

	sli := make([]int, 5, 6)
	for i := 0; i < 5; i++ {
		sli[i] = i
	}
	sli = append(sli, 10)
	sli = append(sli, 10)
	sli = append(sli, 10)

	// fmt.Println("val ", sli)
	// fmt.Println("length ", len(sli))
	// fmt.Println("capacity ", cap(sli))

	// maps

	pirates := make(map[int]string)

	pirates[0] = "Hello world!"
	pirates[2] = "Ok bye, see you tomorrow"

	for k, v := range pirates {
		if k == 1 {
			fmt.Println("Print 1")
			break
		} else {
			fmt.Println("Okkkkk", v)
			break
		}
	}

	prothom, ditio := returnTwo(1, 2)
	fmt.Println(prothom, ditio)
}

func returnTwo(a, b int) (jogfol, biyoufol int) {
	jogfol = a + b
	biyoufol = a - b
	return
}
