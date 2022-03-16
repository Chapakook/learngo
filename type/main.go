package main

import "fmt"

/*
int
string
struct
map[type]type
var
*/

//struct
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
)

func main() {
	// map new two way
	// myVarMap := map[string]string{}
	myVarMap := make(map[string]string)

	myVarMap["first"] = "first item"

	mystruct := myStruct{
		myInt: myVarInt,
		myStr: myVarStr,
		myMap: myVarMap,
	}

	fmt.Println("Struct : ", mystruct)
	fmt.Println("Int : ", mystruct.myInt)
	fmt.Println("String : ", mystruct.myStr)
	fmt.Println("Map : ", mystruct.myMap)
}
