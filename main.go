package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ReplaceAll(input, " ", ""))

	num1, operator, num2, err := parseInput(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	res, err := calculate(num1, operator, num2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", res)
}

func parseInput(input string) (int, string, int, error) {
	operators := []string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if strings.Contains(input, operator) {
			nums := strings.Split(input, operator)
			num1, err := parseNum(nums[0])
			if err != nil {
				return 0, "", 0, err
			}

			num2, err := parseNum(nums[1])
			if err != nil {
				return 0, "", 0, err
			}
			return num1, operator, num2, nil
		}
	}
	return 0, "", 0, fmt.Errorf("неккорректное выражение")
}

func parseNum(numStr string) (int, error) {
	romans := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for indx, roman := range romans {
		if strings.Contains(numStr, roman) {
			num := indx + 1
			if num < 1 || num > 10 {
				return 0, fmt.Errorf("число должно быть от 1 до 10")
			}
			return num, nil
		} else {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, fmt.Errorf("неккорректное число: %s", numStr)
			}

			if num < 1 || num > 10 {
				return 0, fmt.Errorf("число должно быть от 1 до 10")
			}
			return num, nil
		}
	}
	return 0, fmt.Errorf("неккорректное число: %s", numStr)
}

func calculate(num1 int, operator string, num2 int) (int, error) {
	var res int
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 * num2
	default:
		return 0, fmt.Errorf("неккорректный оператор: %s", operator)
	}
	return res, nil
}
