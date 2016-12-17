package basics

import (
	"fmt"
	"math"
)

func init()  {
	fmt.Println("imports package init")
}
// Imports is ...
func Imports(x float64) {
	fmt.Println(math.Sqrt(x))
}
