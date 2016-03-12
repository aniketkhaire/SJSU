package main

import "testing"

func Test_fibonacci(t *testing.T){
	var output int
	output = fibonacci_func(5)
	if output != 5{
		t.Error("Fibonacci Test Failed.");
	}
}