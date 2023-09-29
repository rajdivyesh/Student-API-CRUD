package main

import "fmt"

func main2() {

	// 1.  Define an array and add a string to it, then print the content of this array in the reverse format
	var studentArray [4]int
	studentArray[0] = 10
	studentArray[1] = 12
	studentArray[2] = 18
	studentArray[3] = 25

	// Print the content of the array in reverse order
	for i := len(studentArray) - 1; i >= 0; i-- {
		fmt.Println(studentArray[i])
	}

	// 2. Implement a function which counts the total number of characters in an array
	/*var charArray [5]string
	charArray[0] = "Divyesh"
	charArray[1] = "Raj"

	totalChar := 0
	for i := 0; i < len(charArray); i++ {
		totalChar += len(charArray[i])
	}
	fmt.Printf("Total characters in array: %d\n", totalChar) */

	myArray := []string{"LaSalle", "Thiago", "Simran", "Montreal"}
	/*var myArray []string
	myArray[0] = "LaSalle"
	myArray[1] = "Thiago"
	myArray[2] = "Simran"
	myArray[3] = "Montreal"*/
	total_number := CountCharacters(myArray)
	fmt.Printf("The total number is: %d\n", total_number)

	// 3.implement a code with different functions and define an array into each of them, then swap their contents
	myArray1 := []string{"LaSalle", "Thiago", "Simran", "Montreal"}
	myArray2 := []string{"Divyesh", "raj", "Buddy", "India"}

	fmt.Print("Array1: \n", myArray1)
	fmt.Print("Array2: \n", myArray2)

	myArray1, myArray2 = swapArray(myArray1, myArray2)
	fmt.Print("\nAfter swapping:")
	fmt.Print("Array1: \n", myArray1)
	fmt.Print("Array2: \n", myArray2)
}
func CountCharacters(array []string) int {
	total := 0
	for _, str := range array {
		total += len(str)
	}
	return total
}
func swapArray(arr1, arr2 []string) ([]string, []string) {
	return arr2, arr1
}
