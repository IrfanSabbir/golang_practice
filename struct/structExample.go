package structExample

import (
	"fmt"
	"log"
)

type Person struct {
	name      string
	age       float32
	profetion string
	address   string
}

type Doctor struct {
	Person
	chamber []string
}

func StructExample() {
	fmt.Println("Struct Related Examples.................")
	aDoctor := Doctor{
		Person: Person{
			name:    "Abdul Based",
			address: "Chittagiong",
		},
		chamber: []string{"Kulshi", "probortok"},
	}
	fmt.Println(aDoctor)

	bDoctor := new(Doctor)
	bDoctor.name = "Abdul karim"
	bDoctor.address = "dhaka"
	bDoctor.age = 34
	bDoctor.profetion = "Doc"

	bDoctor.chamber = []string{"BADDA", "MOHAKALI"}

	fmt.Println(bDoctor)

	var cDoctor = Doctor{Person{"Abdul hakim", 12, "sdf", "KUlna"}, []string{"a", "b"}}
	fmt.Println(cDoctor)

	var dDoctor Doctor
	dDoctor.name = "lsakdl"
	log.Println("Log testing")
	log.Fatalln("fatal message")
	return
}
