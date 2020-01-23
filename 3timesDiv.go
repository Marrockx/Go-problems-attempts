// Program to print out all numbers divisible by 3
package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}

	}
}
