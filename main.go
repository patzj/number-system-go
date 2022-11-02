package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	printTitle()

	for {
		printOptions()

		choice := 0
		promptConversion(&choice)

		switch choice {
		case 0:
			goto End
		case 1:
			decimalToBinary()
		case 2:
			decimalToOctal()
		case 3:
			decimalToHexadecimal()
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println()
	}

End:
	fmt.Println("Goodbye!")
	os.Exit(0)
}

func printTitle() {
	fmt.Println("*******************************")
	fmt.Println("|   Number System Converter   |")
	fmt.Println("*******************************")
}

func printOptions() {
	fmt.Println("0 - Exit")
	fmt.Println("1 - Decimal to Binary")
	fmt.Println("2 - Decimal to Octal")
	fmt.Println("3 - Decimal to Hexadecimal")
}

func promptConversion(choice *int) {
	fmt.Print("Enter choice: ")
	if _, err := fmt.Scanf("%d", choice); err != nil {
		os.Exit(1)
	}
}

func promptDecimal(number *int) {
	fmt.Print("Enter decimal number: ")
	if _, err := fmt.Scanf("%d", number); err != nil {
		os.Exit(1)
	}
}

func getDigits(number, base int) (digits []string) {
	for number > 0 {
		rem := toHexDigit(number % base)
		digits = append([]string{rem}, digits...)
		number /= base
	}
	return
}

func decimalToBinary() {
	number := 0
	promptDecimal(&number)

	digits := getDigits(number, 2)
	digits = append([]string{"0", "b"}, digits...) // prefix

	fmt.Println("Binary number:", strings.Join(digits, ""))
}

func decimalToOctal() {
	number := 0
	promptDecimal(&number)

	digits := getDigits(number, 8)
	digits = append([]string{"0", "o"}, digits...) // prefix

	fmt.Println("Hexadecimal number:", strings.Join(digits, ""))
}

func decimalToHexadecimal() {
	number := 0
	promptDecimal(&number)

	digits := getDigits(number, 16)
	digits = append([]string{"0", "x"}, digits...) // prefix

	fmt.Println("Hexadecimal number:", strings.Join(digits, ""))
}

func toHexDigit(digit int) string {
	switch digit {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return fmt.Sprint(digit)
	}
}
