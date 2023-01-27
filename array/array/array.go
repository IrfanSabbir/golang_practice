package array

import (
	"fmt"
	"reflect"
	// "reflect"
)

func Array() {
	fmt.Println("Array Related Examples.................")
	// https://www.golangprograms.com/go-language/arrays.html
	fmt.Println("Find this url helpful: https://www.golangprograms.com/go-language/arrays.html")
	//some ways of array
	var arr1 [5]int

	fmt.Println("Array Example1.................")

	for i := 0; i < len(arr1); i++ {
		arr1[i] = (i + 1) * 10
	}

	arr1[0] = -1

	for _, value := range arr1 {
		fmt.Println(value)
	}
	fmt.Println("Array Example2.................")

	// var arr2 [5]int = [5]int{10, 20, 30}
	// arr2 := [...]int{10, 20, 30}
	arr2 := [5]int{10, 20, 30, 40, 50}

	for _, value := range arr2 {
		fmt.Println(value)
	}

	fmt.Println("Array Example2.................")

	// var arr2 [5]int = [5]int{10, 20, 30}
	// arr2 := [...]int{10, 20, 30}
	arrStr := [5]string{"i", "am", "good"}

	for _, value := range arrStr {
		fmt.Println(value)
	}

	fmt.Println("Array Example3.................")
	arrstr1 := [5]string{"Canada", "asd", "aksdj"}
	arrstr2 := arrstr1
	arrstr3 := &arrstr1

	arrstr1[1] = "bd"

	fmt.Println(arrstr1)
	fmt.Println(arrstr2)
	fmt.Println(*arrstr3)
	for _, value := range *arrstr3 {

		if reflect.ValueOf(value).Kind() == reflect.String && len(value) != 0 {
			fmt.Println(value, " ", reflect.ValueOf(value).Kind())
		} else {
			// panic("Invalid data-type")
			break
		}
	}

	arr := reflect.ValueOf((arrstr1))

	if arr.Kind() == reflect.Array {
		fmt.Println("Item is of type", reflect.Array)
	} else {
		fmt.Println("Item not of type", reflect.Array)
	}

	fmt.Println("Array Example filters.................")
	arrstrF := [5]string{"Canada", "manada", "janada", "lanada", "danada"}
	fmt.Println(arrstrF[1:5])
	fmt.Println(arrstrF[2:])
	fmt.Println(arrstrF[:3])
	fmt.Println(arrstrF[:])

	return
}
