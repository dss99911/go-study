package syntex

import (
	"fmt"
	"math"
)

func SampleIf(a int) {
	if a > 1 {
		fmt.Println("a>1")
	} else if a == 0 {
		fmt.Println("a == 0")
	} else {
		fmt.Println("a < 0")
	}

	//if with declaration
	if v := math.Pow(10, 10); v < 14 {
		fmt.Println(v)
	}

	v := &person{}
	if v != nil {
		fmt.Println("not nil")
	}
}

func SampleIf2() {
	x := SampleMap2()

	fmt.Println(x["first"])
	if v, ok := x["second"]; ok {
		fmt.Println(v)
	}

}
