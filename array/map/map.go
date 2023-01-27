package arraymap

import (
	"fmt"
)

func ArrayMap() {
	fmt.Println("Array Map Related Examples.................")

	countryPop := map[string]int{
		"bangladesh": 12023213,
		"india":      234234234,
		"usa":        234234234,
	}
	countryPop["pakishtan"] = 98249384
	countryPopCopy := countryPop
	delete(countryPopCopy, "india")
	fmt.Println(countryPop)
	fmt.Println(countryPopCopy)

	value, isExist := countryPop["india"]
	_, ok := countryPop["pakishtan"]

	fmt.Println(value, " ", isExist)
	fmt.Println(ok)

	var mapKeyValue = make(map[string]int)

	for index, value := range countryPop {
		fmt.Println(index, value)
		mapKeyValue[index] = value
	}

	fmt.Println(mapKeyValue)
	delete(mapKeyValue, "pakishtan")
	fmt.Println(mapKeyValue)
	fmt.Println(countryPop)

}
