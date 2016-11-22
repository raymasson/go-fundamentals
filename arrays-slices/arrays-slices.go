package main

import (
	"fmt"
)

func main() {
	//Slice
	//myCourses := make([]string, 5, 10)

	//myCourses := []string{"Docker", "Puppet", "Python"}

	//fmt.Printf("Length is: %d \nCapacity is: %d", len(myCourses), cap(myCourses))

	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(mySlice[4])

	mySlice[1] = 0

	fmt.Println(mySlice)

	sliceOfSlice := mySlice[2:5]

	fmt.Println(sliceOfSlice)

	sliceOfSlice2 := mySlice[:5]

	fmt.Println(sliceOfSlice2)

	sliceOfSlice3 := mySlice[2:]

	fmt.Println(sliceOfSlice3)

	myAppendSlice := make([]int, 1, 4)

	fmt.Printf("Length is: %d Capacity is: %d", len(myAppendSlice), cap(myAppendSlice))

	for i := 1; i < 17; i++ {
		myAppendSlice = append(myAppendSlice, i)
		fmt.Printf("\nCapacity is: %d", cap(myAppendSlice))
	}

	mySlice2 := []int{1, 2, 3, 4, 5}
	fmt.Println("\n", mySlice2)

	for _, i := range mySlice2 {
		fmt.Println("for range loop:", i)
	}

	newSlice := []int{10, 20, 30}
	mySlice2 = append(mySlice2, newSlice...)

	fmt.Println("\n", mySlice2)
}
