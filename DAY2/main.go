package main

import (
	"fmt"
)

func main() {
	var name string = "GoLang";
	var age,rollno int = 20,101;

	//multi var declaration
	var (
		a = 5
		b = 10
		c = 15
	)

	//short hand declaration
	var1, var2 := 1,2;

	const pi float32 = 3.14;

	var bvalue bool = true;

	fmt.Print(a,b,c);
	fmt.Print("\n");
	fmt.Println("var1: ",var1,"var2: ",var2);
	fmt.Print("Name: ",name," Age: ",age," Rollno: ",rollno);
	fmt.Println("Value of pi: ",pi);
	fmt.Println("Boolean value: ",bvalue);
	
	//type conversion
	var i int = 230;
	var f float64 = float64(i);
	fmt.Println(f);

	var j float64 = 750.514;
	var k int = int(j);
	fmt.Println(k);
}