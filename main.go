package main

import (
	"fmt"

	linkedListExample "go.practice.com/linkedlist"
)

// all uncometed packages
// array "go.practice.com/array/array"
// arraymap "go.practice.com/array/map"
// arraySlice "go.practice.com/array/slice"
// structExample "go.practice.com/struct"
// deferExamp "go.practice.com//deferPanicRecover/defer"
// panicRecoverExample "go.practice.com/deferPanicRecover/panicRecover"
// interfaceExample "go.practice.com/interface"
// goroutineExample "go.practice.com/goroutine"
// linkedListExample "go.practice.com/linkedlist"

func main() {
	fmt.Println("Go practice environment......................using go 1.19")

	//Array related topic
	// Uncomment below line
	// array.Array()

	//Array Slice related topic
	// Uncomment below line
	// arraySlice.ArraySlice()

	//Array Map related topic
	// Uncomment below line
	// arraymap.ArrayMap()

	//Struct related topic
	// Uncomment below line
	// structExample.StructExample()

	// Defer related topic
	// Uncomment below line
	// deferExamp.DeferExample()

	// Panic Recover related topic
	// Uncomment below line
	// panicRecoverExample.PanicRecoverExample()

	// Interface Recover related topic
	// Uncomment below line
	// interfaceExample.InterfaceExample()

	// Goroutine Recover related topic
	// Uncomment below line
	// goroutineExample.GoroutineExample()

	// Goroutine Recover related topic
	// Uncomment below line
	linkedListExample.LinkedListExample()

	fmt.Println("Go practice endded")
}
