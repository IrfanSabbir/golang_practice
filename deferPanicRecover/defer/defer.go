package deferExamp

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeferExample() {
	log.Println()
	fmt.Println("Defer Related Examples.................\n")

	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	log.Println()
	return
}
