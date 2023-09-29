package main

import "fmt"

/*
C++: int min(){}
Java and C#: public static void main(String[] args){}
Python: main()
Golang: func main()
*/

/*
C#/C++/JAVA : public [RETURN TYPE] [NAME OF THE METHOD](ARGUMENTS)
private string name(int a, int b)
public void add(int a, int b)
*/

// in the below function the return type is int
func add(a, b int) int {
	return a + b
}

// since in below function/method we have not defined any return type, then by default it is void
func myInfo(name string, age int) {
	fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}
func main1() {
	result := add(10, 19)
	fmt.Printf("The sum value is: %d \n", result)

	myInfo("Divyesh", 23)
}
