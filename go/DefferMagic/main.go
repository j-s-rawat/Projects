package main

import "fmt"

var i int

func testingDefer() (j *int) {
	defer func() {
		i = 1 //the global i is incremented to 1

		j = &i //local j is initialized in stack but is excapes to heap as
		//it is being returned and the life of this variable will be beyong the scope of function
	}()
	return j
}

func testingDefer2() (i int) {
	defer func() {
		fmt.Println("in defer value of i:", i) //as now 'i' is added as named return variable, in this function 'i' is now a local function variable
		i = 2
		fmt.Println("in defer value of i:", i)
	}()
	return i
}
func main() {
	fmt.Println("function return value: ", testingDefer()) //even though i is being modified in the deffer, 'i' is being modified only when function is removed from stack
	fmt.Println(i)

	fmt.Println("function return value: ", testingDefer2())
	fmt.Println(i)

}
