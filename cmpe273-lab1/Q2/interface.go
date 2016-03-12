package main

import "fmt"

type Shape interface{
	perimeter() float64
}

type Rectangle struct{
	length, breadth float64
}

func (rect Rectangle)Area() float64{
	return rect.length*rect.breadth
}
func (rect Rectangle)perimeter() float64{
	return (rect.length+rect.breadth)*2
}

type Circle struct{
	radius float64
}

func (circ Circle)Area() float64{
	return circ.radius*circ.radius*3.14
}
func (circ Circle)perimeter() float64{
	return 2*circ.radius*3.14
}


func main(){
	rect := Rectangle{0,0}
	fmt.Println("Enter length and breadth of Rectangle: ")
	fmt.Scanf("%f %f", &rect.length, &rect.breadth)

	var s Shape
	s = rect
	fmt.Println("Rect Area = ", rect.Area())
	fmt.Println("Rect Perimeter = ", s.perimeter())

	circ := Circle{10}
	/*fmt.Println("Enter radius of Circle: ")
	fmt.Scanln( &circ.radius)
	*/
	s = circ
	fmt.Println("Circle Area = ", circ.Area())
	fmt.Println("Circle Perimeter = ", circ.perimeter())
}