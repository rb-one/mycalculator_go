package mycalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

func (c calc) operate(input string, operator string) int {

	cleanInput := strings.Split(input, operator)
	number1 := parsear(cleanInput[0])
	number2 := parsear(cleanInput[1])

	switch operator {

	case "+":
		fmt.Println(number1 + number2)
		return (number1 + number2)
	case "-":
		fmt.Println(number1 - number2)
		return (number1 - number2)
	case "*":
		fmt.Println(number1 * number2)
		return (number1 * number2)
	case "/":
		fmt.Println(number1 / number2)
		return (number1 / number2)
	default:
		fmt.Println(operator, "it's not supported")
		return 0
	}
}

func parsear(input string) int {
	number, _ := strconv.Atoi(input)
	return number
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {

	input := readInput()
	operator := readInput()

	fmt.Println(input, operator)

	c := calc{}
	fmt.Println()
	c.operate(input, operator)

}
