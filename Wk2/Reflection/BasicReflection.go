package main

//import
import (
	"fmt"
	"reflect"
)

// define variables
var (
	a = "This is a string"
	b = 12345
	c = 1.2345
	d = true
)

// create interface
type empty interface{}

// print out the reflected type and value
func inspect1(e empty) {
	fmt.Println(reflect.TypeOf(e), reflect.ValueOf(e))
}

// main function
func main() {
	//activity 1
	inspect1(a)
	inspect1(b)
	inspect1(c)
	inspect1(d)

	//activity2
	customer1 := customer{
		FName:        "John",
		LName:        "Wick",
		UserID:       123123123,
		InvoiceTotal: 10000,
	}
	inspect2(customer1)
}

// create customer structure
type customer struct {
	FName        string
	LName        string
	UserID       int
	InvoiceTotal float64
}

func inspect2(n interface{}) {

	refType := reflect.TypeOf(n)
	refValue := reflect.ValueOf(n)
	fmt.Println("No. of fields: ", refType.NumField())
	fmt.Println(refType.Name())
	fmt.Println(refType.Kind())
	for i := 0; i < refType.NumField(); i++ {
		fmt.Println(refType.Field(i).Name, "value: ", refValue.Field(i), "type: ", refType.Field(i).Type)
	}

}
