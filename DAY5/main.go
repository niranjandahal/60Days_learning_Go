package main

import "fmt"

func main() {

    var arr [5]int
    arr[0] = 1; arr[1] = 2; arr[2] = 3; arr[3] = 4; arr[4] = 5;
    fmt.Println("Array:", arr)

    arr2 := [3]string{"Go", "is", "fun"}
    fmt.Println("Array2:", arr2)

    // Slices
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println("Slice:", slice)

    slice = append(slice, 6, 7)
    fmt.Println("Slice after appending:", slice)

    subSlice := slice[2:5]
    fmt.Println("Sub-slice:", subSlice)

	//print with indexing
    for i, v := range slice {
        fmt.Printf("Index: %d, Value: %d\n", i, v)
    }
	//print without indexing
    for _, v := range slice {
        fmt.Printf("Value: %d\n", v)
    }

    fmt.Printf("Length of slice: %d\n", len(slice))
    fmt.Printf("Capacity of slice: %d\n", cap(slice))

    //slices using mak
    madeSlice := make([]int, 5, 10)
    fmt.Printf("Made Slice: %v, Length: %d, Capacity: %d\n", madeSlice, len(madeSlice), cap(madeSlice))

    //slices copy
    srcSlice := []int{1, 2, 3}
    destSlice := make([]int, len(srcSlice))
    copy(destSlice, srcSlice)
    fmt.Println("Source Slice:", srcSlice)
    fmt.Println("Destination Slice after copy:", destSlice)
}