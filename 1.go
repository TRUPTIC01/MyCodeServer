package main

import "fmt"

func main(){
	adder(add)
	adder1(add)
}

type myFunctionDataType func (int, int ) int // data type and variable of function (simplifying)
			//name  //data type
func adder1 ( add myFunctionDataType) { //function is given as a parameter input
	val := add(3,5)
	fmt.Println(val)
}


func adder ( add func (int, int ) int ) { //function is given as a parameter input
	val := add(3,5)
	fmt.Println(val)
}
func add(i int, j int) int{
	return i+j
}