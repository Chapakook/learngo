package main

import "fmt"

/*
int
string
struct
map[type]type
var
*/

type myStruct struct {
	myInt int
	myStr string
	myMap map[string]string
}

// var use two way
var myvar = ""

// same var combine ()  (ex. error ...)
var (
	myVarInt = 0
	myVarStr = "string"
	myVarMap = make(map[string]string)
)

func main() {
	myVarMap["first"] = "first item"

	mystruct := myStruct{
		myInt: myVarInt,
		myStr: myVarStr,
		myMap: myVarMap,
	}

	fmt.Println(mystruct)
	fmt.Println(mystruct.myInt)
	fmt.Println(mystruct.myStr)
	fmt.Println(mystruct.myMap)
}
