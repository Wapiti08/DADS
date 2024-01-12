package main

import (
	"fmt"
	"reflect"
)

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}

	mys := myStruct{2, "Hello", 2.4}
	InspectStructType(&mys)

}

func InspectStructType(i interface{}) {
	// the interface can be any type of data
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr {
		return
	}
	// Elem returns the types' element type
	mysRValue = mysRValue.Elem()
	// kind returns the type of value
	if mysRValue.Kind() != reflect.Struct {
		return
	}

	mysRValue.Field(0).SetInt(15)

	mysRType := mysRValue.Type()

	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i)
		fieldRValue := mysRValue.Field(i)

		fmt.Println("FieldName: '%s', field type: '%s', field value: '%v' \n",
			fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
		fmt.Println("Struct tags, alias: ", fieldRType.Tag.Get("alias"), fieldRType.Tag.Get("desc"))
	}
}
