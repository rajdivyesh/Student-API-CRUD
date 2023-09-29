package main

import "fmt"

func main22() {
	i, o := []int{12, 0, 9, 13, 7, 52, 8}, make([]bool, len([]int{12, 0, 9, 13, 7, 52, 8}))
	for j, n := range i {
		o[j] = func() bool {
			if n <= 1 {
				return false
			}
			if n <= 3 {
				return true
			}
			for i := 2; i*i <= n; i++ {
				if n%i == 0 {
					return false
				}
			}
			return true
		}()
	}
	fmt.Println(o)
}
