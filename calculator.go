package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumerals = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var integersToRomans = []struct {
	value int
	symbol string
}{
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic("Ошибка чтения ввода")
	}

	result, err := calculate(strings.TrimSpace(input))
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(result)
}

func calculate(input string) (string, error) {
	tokens := tokenize(input)
	if len(tokens) != 3 {
		return "", fmt.Errorf("некорректный формат математической операции")
	}

	aStr, operator, bStr := tokens[0], tokens[1], tokens[2]

	isRoman := isRomanNumeral(aStr) && isRomanNumeral(bStr)
	isArabic := isArabicNumeral(aStr) && isArabicNumeral(bStr)

	if !isRoman && !isArabic {
		panic("используются одновременно разные системы счисления")
	}

	if isRoman {
		a, _ := romanToInt(aStr)
		b, _ := romanToInt(bStr)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			return "", fmt.Errorf("римские числа должны быть от I до X")
		}
		result, err := performOperation(a, b, operator)
		if err != nil {
			panic(err.Error())
		}
		if result < 1 {
			panic("в римской системе нет отрицательных чисел или нуля")
		}
		return intToRoman(result), nil
	}

	if isArabic {
		a, _ := strconv.Atoi(aStr)
		b, _ := strconv.Atoi(bStr)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("числа должны быть от 1 до 10")
		}
		result, err := performOperation(a, b, operator)
		if err != nil {
			panic(err.Error())
		}
		return strconv.Itoa(result), nil
	}

	panic("используются одновременно разные системы счисления")
}

func tokenize(input string) []string {
	return strings.Fields(input)
}

func isRomanNumeral(s string) bool {
	_, exists := romanNumerals[s]
	return exists
}

func isArabicNumeral(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func romanToInt(s string) (int, error) {
	if val, exists := romanNumerals[s]; exists {
		return val, nil
	}
	panic("некорректное римское число")
}

func intToRoman(num int) string {
	var result strings.Builder
	for _, roman := range integersToRomans {
		for num >= roman.value {
			result.WriteString(roman.symbol)
			num -= roman.value
		}
	}
	return result.String()
}

func performOperation(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			panic("деление на ноль")
		}
		return a / b, nil
	default:
		panic("некорректный оператор")
	}
}
