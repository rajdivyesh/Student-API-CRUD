package main

import "fmt"

func main3() {
	/*
		// Address of a variable: &
		// Value of the variable: *

		// The below code means, we have a local variable which its name is num, its type is integer and its value is 20

		var num int = 20

		// The below line means, we have a pointer to an integer variable but we have no value for it

		var ptr *int

		ptr = &num

		fmt.Println("The value of num: ", num)
		fmt.Println("The value of the location where ptr is referring to: ", *ptr)

		*ptr = 56
		fmt.Println("The new value of the pointer is: ", *ptr)

	*/
	// Implement a code as below:
	/*
					Define two variables with any type and print their address
					Is there any visual difference between the address of an integer variable and string?


				var num2 int = 22
				var name string = "test123"

				var ptr1 *int
				var name1 *string

				ptr1 = &num2
				name1 = &name

				*ptr1 = 66

				fmt.Println("The value of num: ", num2)
				fmt.Println("The value of the location where ptr1 is referring to: ", *ptr1)
				fmt.Println("The  location where ptr1 is referring is: ", ptr1)

				fmt.Println("The value of string: ", name)
				fmt.Println("The value of the location where name1 is referring to: ", *name1)
				fmt.Println("The location where name1 is stored is: ", name1)




			// Below lines defines an array of three elements with pointer
			var ptrArray [3]*int

			// below lines are for defining three local variable without pointers
			num1 := 34
			num2 := 18
			num3 := 98

			// In the below line we assign the address pof num1 to the first location of the array
			ptrArray[0] = &num1
			// In the below line we assign the address pof num2 to the second location of the array
			ptrArray[1] = &num2
			// In the below line we assign the address pof num3 to the third location of the array
			ptrArray[2] = &num3

			fmt.Println("Value at the first array location: ", *ptrArray[0])
			fmt.Println("Value at the second array location: ", *ptrArray[1])
			fmt.Println("Value at the third array location: ", *ptrArray[2])


		 //below works with the func findmymin
		myArray := []int{45, 3, 12, 2, 86, 4, 9, 1, 25}
		result := findMyMin(myArray)
		if result != nil {
			fmt.Printf("The minimum value of your array is: %d", *result)
		} else {
			fmt.Println("Your array is empty!")
		}

	*/

	myArray1 := []float32{12, 45, 5, 9, 18, 76, 2, 20}
	resultMax := findMyMax(myArray1)
	resultMin := findMyMin(myArray1)
	if resultMax != nil && resultMin != nil {
		fmt.Println("Maximum value in the array: ", *resultMax)
		fmt.Println("The minimum value of your array is: ", *resultMin)
	} else {
		fmt.Println("The array is empty.")
	}

}
func findMyMax(array []float32) *float32 {
	if len(array) == 0 {
		return nil
		// recall that nil is the same as null in other programming languages

	}
	max := &array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > *max {
			max = &array[i]
		}
	}
	return max
}

func findMyMin(array []float32) *float32 {
	if len(array) == 0 {
		return nil
		// recall that nil is the same as null in other programming languages

	}
	min := &array[0]
	for i := 1; i < len(array); i++ {
		if array[i] < *min {
			min = &array[i]
		}
	}
	return min
}
