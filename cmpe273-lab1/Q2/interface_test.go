package main

import "testing"

func Test_interface(t *testing.T){
	
	rect := Rectangle{10,10}
	circ := Circle{10}
	var s Shape
	var result float64

	s = rect
	result = s.perimeter()
	if result != 40{
		t.Error("Error in Rectangle Perimeter")
	}

	s = circ
	result = s.perimeter()
	if result != 62.800000000000004{
		t.Error("Error in Circle Perimeter")
	}
}