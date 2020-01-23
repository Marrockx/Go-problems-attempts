package man

import "fmt"

func main() {
	// define variables

	var lengthInFeet, lengthinMetre float32

	fmt.Println("What is the length measured in ft? ")
	fmt.Scanln(&lengthInFeet)

	// Note that 1 feet(ft) = 0.3048 metre(m)
	lengthinMetre = (lengthInFeet * 0.3048)
	fmt.Println("The length in metre(m) will be: ", lengthinMetre, "m")
}
