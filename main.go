package main

import "os"

// atoi converts a string to an integer.
func atoi(s string) int {
	result := 0
	sign := 1
	signDetected := false

	// Iterate over each character in the string.
	for i, char := range s {
		// Handle signs at the beginning of the string.
		if char == '-' || char == '+' {
			if signDetected {
				break
			}
			if i != 0 {
				return 0 // Invalid format: sign detected in the middle of the string.
			}
			if char == '-' {
				sign = -1
				continue
			}
			signDetected = true
			continue
		}

		// Check if the character is a digit.
		if char < '0' || char > '9' {
			return 0 // Invalid character found.
		} else {
			digit := int(char - '0')
			result = result*10 + digit
		}
	}

	return result * sign
}

// itoa converts an integer to a string.
func itoa(num int) string {
	if num == 0 {
		return "0"
	}
	sign := ""
	if num < 0 {
		sign = "-"
		num = -num
	}

	result := ""

	// Extract digits from the number and build the result string.
	for num > 0 {
		digit := num % 10
		result = string('0'+rune(digit)) + result
		num /= 10
	}

	return sign + result
}

func main() {
	// Check if the program is called with the correct number of arguments.
	if len(os.Args) != 4 {
		os.Stdout.WriteString("Usage:<value1><operator><value2>\n")
		return
	}

	// Parse command-line arguments.
	value1 := atoi(os.Args[1])
	operator := os.Args[2]
	value2 := atoi(os.Args[3])

	result := 0

	// Perform arithmetic operation based on the operator.
	switch operator {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		// Handle division by zero case.
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		} else {
			result = value1 / value2
		}
	case "%":
		// Handle modulo by zero case.
		if value2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		} else {
			result = value1 % value2
		}
	default:
		os.Stdout.WriteString("Invalid operator\n")
	}

	// Convert the result to a string and write it to standard output.
	output := itoa(result)
	os.Stdout.WriteString(output)
	os.Stdout.WriteString("\n")
}
