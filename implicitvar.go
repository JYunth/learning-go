package main

import "fmt"

func main(){
	var basketTotal, totalPrice float64 //Multiple variable declaration
	carrotPrice := 0.75 // The compiler analyzes the value and then makes the variable type the most relevant type
	basketTotal = basketTotal + carrotPrice
	fmt.Println(basketTotal)

	// Output is 0.75 because basketTotal is declared and not assigned any value.
	// Which means that the value of basketTotal before addition was 0

	spinachPrice := 1.50
	basketTotal += spinachPrice
	// basketTotal = basketTotal + spinachPrice // The above can also be written like this
	fmt.Println(basketTotal)

	// Output is 2.25 which is the previous output (0.75) + spinachPrice (1.5)
	basketTotal += totalPrice
	
	command := "hold my "
	beverage := "soda"
	command += beverage //Concatenation operator
	fmt.Println(command)

	// Prints hold my soda
}