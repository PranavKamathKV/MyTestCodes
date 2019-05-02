package main

import (
	"fmt"
	"reflect"
)

type Person struct{
	age int
	name string
}

type Address struct{
	housename string
	streetnum int
	city string
	landmark string
}


func useReflection(example interface{}){
	// Add more reflection types here
	mySpecificType:= reflect.TypeOf(example).Kind()
	if mySpecificType == reflect.Struct {
		myType:=reflect.TypeOf(example).Name()
		myValues:= reflect.ValueOf(example)

		fmt.Println("Type, Value: ",myType, myValues) 

		//Use numfield method to get the number of fields in the myValues
		for i:=0; i<myValues.NumField();i++{
			// Demo to use Field method with Kind()
			switch myValues.Field(i).Kind() {
			case reflect.Int:
				fmt.Println("Kind int in struct", myValues.Field(i))
			case reflect.String:
				fmt.Println("Kind string in struct",myValues.Field(i))
			default:
				fmt.Println("Not supported")

			}
		}

	}else{
		fmt.Println("Unsupported type %T",example)
		fmt.Println("Value: ", example)
	}

	return

}

func main(){
	jack:= Person{
		age: 25,
		name: "Jack Crow",
	}

	useReflection(jack)

	jackAddr:= Address{
		housename: "Green Villa",
		streetnum: 24,
		city: "Tacoma",
		landmark: "Chihuly Bridge of Glass",
	}

	useReflection(jackAddr)

	// some more types to handle at runtime
	num:= 800
	name:="Reflection"
	useReflection(num)
	useReflection(name)
}
