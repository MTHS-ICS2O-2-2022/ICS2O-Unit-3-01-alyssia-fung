// Created by: Alyssia fung
// Created on: Mar 2023
//
// This program finds the area of a trapezoid

package main

import "fmt"

func main() {
	// Define the variables for the trapezoid's bases and height
	var side1, side2, height float64

	// input
	fmt.Println("This program finds the area of a trapezoid")
	fmt.Println()
	fmt.Print("Enter the length of side1 (cm): ")
	fmt.Scanln(&side1)
	fmt.Print("Enter the length of side2 (cm): ")
	fmt.Scanln(&side2)
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)

	// process
	area := (side1 + side2) * height / 2

	// print out the calculated area of the trapezoid
	fmt.Printf("The area of the trapezoid is: %.2f\n", area)
	fmt.Println("\nDone.")
}
