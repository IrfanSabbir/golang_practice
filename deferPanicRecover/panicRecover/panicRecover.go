package panicRecoverExample

import (
	"errors"
	"fmt"
	"log"
)

func PanicRecoverExample() {
	log.Println()
	fmt.Println("Panic recover Examples.................\n")

	myPPanicFunc(1)
	myPPanicFunc(0)
	myPPanicFunc(8)

	log.Println()
	return
}

func myPPanicFunc(i int) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	if i == 0 {
		log.Println("Stopping/....")
		panic(errors.New("I am provided wrong value"))
	} else {
		fmt.Println(i)
	}
}
