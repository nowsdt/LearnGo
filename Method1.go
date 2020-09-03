package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//Point类型的方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 传值函数
func modify(p int) {
	p = 10
}

// 指针函数
func modifyP(p *int) {
	*p = 10
}

// 指针方法
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 结构体内嵌组成类型 方法
// 内嵌方法编译器会额外生成包装方法来调用bridge
type ColoredPoint struct {
	Point
	Color color.RGBA
}

type ColoredPointP struct {
	*Point
	Color color.RGBA
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance())

	i := 15
	modify(i)
	fmt.Println("i:", i)
	modifyP(&i)
	fmt.Println("i:", i)

	fmt.Println(p)
	// 报错
	//&p.ScaleBy(10)
	(&p).ScaleBy(10)
	// 编译器会对变量进行&p的隐式转换
	p.ScaleBy(10)
	fmt.Println(p)

	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var cpRed = ColoredPoint{Point{1, 1}, red}
	var cpBlue = ColoredPoint{Point{1, 1}, blue}

	fmt.Println(cpRed.Distance(cpBlue.Point))

	cpRed.ScaleBy(2)
	cpBlue.ScaleBy(2)

	fmt.Println(cpRed.Distance(cpBlue.Point))

	a := ColoredPointP{&Point{1, 1}, red}
	b := ColoredPointP{&Point{1, 1}, blue}

	a.Distance(*b.Point)

}
