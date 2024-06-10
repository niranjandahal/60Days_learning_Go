package main

import "fmt"

//struct shows some attributes

// type Dog struct {
//     name     string
//     age       int
//     gender  string
//     isHungry bool
// }

//interface isn't concerned with state or attributes but with behaviour

type Dog interface {
    barks() 
    eats()
}


type Labrador struct {
    name     string
    age    int
    gender string
    isHungry bool
}

func addToGroup(d Dog, group []Dog) []Dog {
    group = append(group, d)
    return group
}

func (l Labrador) barks() {
    fmt.Println(l.name + " says woof")
}



func (l Labrador) eats() {
    fmt.Println(l.name + " says woof")

   
}

func main() {
    bigDogs := []Dog{}
    max := Labrador{
        name:     "Max",
        age:      5,
        gender:   "Male",
        isHungry: true,
    }


    fmt.Println("Our group of big dogs:", bigDogs)
    bigDogs = addToGroup(max, bigDogs)
    fmt.Println("Our group of big dogs now:", bigDogs)
}
