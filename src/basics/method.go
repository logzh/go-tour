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

const(
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color Color
}

type BoxList []Box

func (b Box)volume()(float64)  {
	return b.depth * b.height * b.width
}

func (b *Box)setColor(c Color)  {
	b.color = c
}

func (bl BoxList)getBiggestBoxColor()(Color)  {
	v := 0.0
	k := Color(WHITE)

	for _, b := range bl{
		if vol := b.volume(); vol > v{
			v = vol
			k = b.color
		}
	}

	return k
}

func init()  {
	fmt.Println("method的用法：")

	r := Rectangle{2, 6}
	fmt.Println("长方形的面积：", r.area())

	c := Circle{8.9}
	fmt.Println("圆的面积：", c.area())

	boxes := BoxList{Box{1, 2, 3, WHITE}, Box{1, 9, 3, YELLOW}, Box{1.8, 2.9, 3.9, BLACK}}

	boxes[len(boxes)-1].setColor(RED)

	fmt.Println("boxes length:", len(boxes))
	fmt.Println("get a box color:", boxes[len(boxes)-1].color)
	fmt.Println("the max box color:", boxes.getBiggestBoxColor())
}