package main

import "fmt"

func main(){

	fmt.Println("Enter the value of n in fib(n): ")
    	var n int = 0
    	fmt.Scanln(&n)
	
	if n < 0{
		fmt.Println("Please enter a positive number")
	}else{
		var res int = fibonacci_func(n)
    		fmt.Println("fib(",n,") = ", res)
	}
}

func fibonacci_func(n int) int{
	if n ==0{
		return 0
	}else if n==1{
		return 1
	}else{
		return (fibonacci_func(n-1)+fibonacci_func(n-2))
	}
}