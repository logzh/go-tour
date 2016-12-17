package basics

import (
	"math"
	"fmt"
)

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) area()(float64)  {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle)area()(float64)  {
	return c.radius * c.radius*math.Pi
}

func init()  {
	fmt.Println("method的用法：")

	r := Rectangle{2, 6}
	fmt.Println("长方形的面积：", r.area())

	c := Circle{8.9}
	fmt.Println("圆的面积：", c.area())
}