package main

import "fmt"

// 	AUTHOR - MARIAM ADEKOLA
// program to convert from degree Fahrenheit to degree Celsius

func main() {
	var tempFahrenheit, tempCelsius float32

	fmt.Println("Enter temperature in Fahrenheit(F): ")
	fmt.Scanln(&tempFahrenheit)
	tempCelsius = (tempFahrenheit - 32) * (5.0 / 9)
	fmt.Println("The temperature in Celsius(C) is: ", tempCelsius)

}
