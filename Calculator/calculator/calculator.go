package calculator

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input string) (float64, error) {
	tokens := strings.Fields(input)

	if len(tokens) < 3 {
		return 0, errors.New("недостатньо операндів")
	}

	result, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		return 0, err
	}

	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operand, err := strconv.ParseFloat(tokens[i+1], 64)
		if err != nil {
			return 0, err
		}

		switch operator {
		case "+":
			result += operand
		case "-":
			result -= operand
		case "*":
			result *= operand
		case "/":
			if operand == 0 {
				return 0, errors.New("ділення на нуль заборонено")
			}
			result /= operand
		default:
			return 0, fmt.Errorf("невідомий оператор: %s", operator)
		}
	}

	return result, nil
}
