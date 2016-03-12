package main

import "time"
import "testing"

func Test_sleep(t *testing.T){
	t0 := time.Now()
    mySleep(5)
	t1 := time.Now()
	if t1.Sub(t0) <= 5{
		t.Error("Error occured in sleep function")
	}
}