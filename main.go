package main

import (
	"fmt"
	"goproject/conditionals"
	"goproject/functions"

	"reflect"
)

func main() {

	// Variable creation
	// var or const
	// short hand declaration
	// var  variableName type
	// var variableName="string"  // you have to give the initial value

	//var myNum = 9
	//var myNum int = 9 // valid syntax

	myNum := 90 // short hand declaration  automatically detects the type
	fmt.Println(reflect.TypeOf(myNum))
	fmt.Println("hello world", myNum)

	// const variables
	const myvariable int = 900
	fmt.Println("hello world", myvariable)

	const myvariable2 = 900
	fmt.Println("hello world", myvariable2)
	// you cant change the value assigned to a const

	//myvariable2 = 999 // invalid through an error

	// In go we have two kind of data types
	// primitive data type  that are present by default in go (int ,int64 ,float64,string ,bool,uint ,uint64)

	// User defined types
	// Soem user defined types are struct ,interfaces ,and you can default primitive types to create the custom types as well

	// Struct is a composite data type it can contain primitive data types as well and  user defined data types too
	type Address struct {
	}

	type StructName struct {
		Name    string
		Age     int
		Address Address
	}

	// interfaces
	// Interface can only contain method declartion
	type InterfaceName interface {
		Send()
		SendToUSer(name string) string // A method signature is composed of the type of parameter its taking and the type of return value
	}

	// Custom types with default priimitive types

	type Myinteger int

	var c Myinteger // custom type

	c = 9

	fmt.Println(reflect.TypeOf(c))
	// Type conversion to convert your type into the primitive type
	z := int(c) // type conversion
	fmt.Println(reflect.TypeOf(z))
	// You can create a function type as well
	type myfunc func()
	var ck myfunc
	fmt.Println(reflect.TypeOf(ck))

	//functions.Sum()
	cd, ke := functions.WithArgsAndMulReturn("hello people")
	fmt.Println(cd, ke)
	fmt.Println(functions.WithMulArgsAndMulReturn("hh", 99))

	fmt.Println(conditionals.CheckIfEven(88))
	fmt.Println(conditionals.CheckIfEven(889))
	fmt.Println(conditionals.CheckIfEven(-1))

	conditionals.SimpleForLoop()
	conditionals.SimpleForLoopAsWhile()
	conditionals.Infiniteloop()

}
