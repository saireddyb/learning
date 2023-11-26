package main

import "fmt"

type circle struct {
	radius float32 
}

type rectangle struct {
	length float32
	breadth float32
}

type shape interface{
	area() float32
}

func (r rectangle) area() float32 {
	return r.length*r.breadth
}

func (c circle) area() float32 {
	return c.radius * 3.14
}

func isPositive(s shape){
	area := s.area()
	if area > 0 {
		fmt.Println("area is postive")
	}
	if area < 0 {
		fmt.Println("area is negative")
	}
}

func main(){
	circle1 := circle{5}
	isPositive(circle1)
	rectangle1 := rectangle{2,3}
	isPositive(rectangle1)
}