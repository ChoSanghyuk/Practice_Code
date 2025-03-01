package gotime

import (
	"fmt"
	"time"
)

func HowToAddOrSub() {

	// Add
	n1 := time.Now()
	f := n1.Add(time.Duration(10) * time.Minute)
	fmt.Println(f)

	// Sub
	n2 := time.Now()
	b := n2.Add(time.Duration(-10) * time.Minute)
	fmt.Println(b)
}

func Today() {

	today := time.Now().Format("2006-01-02 03")
	fmt.Println(today)
}
