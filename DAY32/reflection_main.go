package main

import (
	"fmt"
	"reflect"
)

type User struct {
    Name string
    Age  int
}

func main() {

	var x float64 = 3.4
    t := reflect.TypeOf(x)
    v := reflect.ValueOf(x)

	fmt.Println("Type:", t)
    fmt.Println("Kind:", t.Kind())
    fmt.Println("Value:", v) 

	z := reflect.ValueOf(&x).Elem()
    z.SetFloat(7.1)
    fmt.Println("Modified Value:", x)

	
	user := User{"Alice", 30}
    a := reflect.ValueOf(&user).Elem()
    a.FieldByName("Name").SetString("Bob")
   	a.FieldByName("Age").SetInt(25)
    fmt.Println(user)


    intStack := Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    fmt.Println(intStack.Pop())

    stringStack := Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")
    fmt.Println(stringStack.Pop()) 
}
