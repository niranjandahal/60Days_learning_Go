package main

import "fmt"

type temp struct {
	Lat, Long float64
}

//using make func to create map
var m = make(map[string]temp)

func main() {
    capitals := make(map[string]string)

    capitals["Nepal"] = "Kathmandu"
    capitals["USA"] = "Washington, D.C."
    capitals["Japan"] = "Tokyo"
    
    fmt.Println("Capitals:", capitals)

    // Update
    capitals["USA"] = "Washington"
    fmt.Println("Updated Capitals:", capitals)

    // DeletE
    delete(capitals, "Japan")
    fmt.Println("After Deleting Japan:", capitals)

    // Access
    capitalOfNepal, exists := capitals["Nepal"]
    if exists {
        fmt.Println("The capital of Nepal is:", capitalOfNepal)
    } else {
        fmt.Println("Capital not found")
    }

    for country, capital := range capitals {
        fmt.Printf("The capital of %s is %s\n", country, capital)
    }

    // Checking if a key exists
    _, exists = capitals["India"]
    if !exists {
        fmt.Println("India is not in the map")
    }

    temperatures := map[int]float64{
        1:  98.6,
        2:  99.5,
        3:  100.4,
    }
    fmt.Println("Temperatures:", temperatures)
    
    // Update
    temperatures[1] = 97.7
    fmt.Println("Updated Temperatures:", temperatures)

    // Delete
    delete(temperatures, 2)
    fmt.Println("After Deleting Key 2:", temperatures)

	// printing values of map		
	m["pokhara"] = temp{
		40.68433, -74.39967,
	}
	fmt.Println(m["pokhara"])
}

