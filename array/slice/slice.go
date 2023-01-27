package arraySlice

import (
	"fmt"
	"reflect"
	"strconv"
)

func ArraySlice() {

	fmt.Println("Array Slice Related Examples.................")
	// https://www.golangprograms.com/go-language/slices-in-golang-programming.html
	fmt.Println("Find this url helpful: https://www.golangprograms.com/go-language/slices-in-golang-programming.html")

	var arrSlice []int

	fmt.Println(reflect.ValueOf(arrSlice).Kind())

	var stringSlice = []string{"japan", "apan", "lapan", "mapan"}
	fmt.Println(reflect.ValueOf(stringSlice).Kind())

	for key, value := range stringSlice {
		fmt.Println(key, " is ", value)
	}

	var intSlice = new([50]string)[0:10] // capacity 50 but len is 10
	for i := 0; i < len(intSlice); i++ {
		intSlice[i] = strconv.Itoa(i)
	}
	fmt.Println(intSlice)
	intSlice = append(intSlice, "10", "11", "12", "13")
	intSlice[10] = "100"
	fmt.Println(intSlice, " current length ", len(intSlice))
	fmt.Println(intSlice[2:9])
	fmt.Println(removeIndex(intSlice, 2), " new length after remove ", len(removeIndex(intSlice, 2)))

	return
}

func removeIndex(strItems []string, index int) []string {
	return append(strItems[0:index], strItems[index+1:]...)
}
