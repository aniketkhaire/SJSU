package main

import "fmt"
import "time"

func mySleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}

func main() {
	fmt.Println("Enter the sleep duration: ")
	var dur int
	fmt.Scanf("%d", &dur)
	fmt.Println("Time before sleep: ", <-time.After(time.Second))
	t0 := time.Now()
    mySleep(dur)
	t1 := time.Now()
	fmt.Printf("Slept for %v seconds.\n", t1.Sub(t0))
	fmt.Println("Time after sleep: ", <-time.After(time.Second))
}