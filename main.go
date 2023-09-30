package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	romans = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите выражение:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ReplaceAll(input, " ", ""))

	nums, operator, err := parseInput(input)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	res, err := parseSystemCount(nums, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Результат:", res)

}

func parseInput(input string) ([]string, string, error) {
	operators := []string{"+", "-", "*", "/"}
	for _, operator := range operators {
		if strings.Contains(input, operator) {
			nums := strings.Split(input, operator)
			if len(nums) == 2 {
				return nums, operator, nil
			} else {
				return []string{}, "", fmt.Errorf("формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
			}

		}
	}
	return []string{}, "", fmt.Errorf("неккорректный оператор")
}

func parseSystemCount(nums []string, operator string) (string, error) {
	var res string
	var err error
	var q int
	for _, num := range nums {
		for _, roman := range romans {
			if num == roman {
				q += 1
			}
		}
	}
	switch q {
	case 0:
		num1, err := parseArabNum(nums[0])
		if err != nil {
			return "", err
		}
		num2, err := parseArabNum(nums[1])
		if err != nil {
			return "", err
		}
		res, err = arabCalc(num1, num2, operator)
		if err != nil {
			return "", err
		}
	case 1:
		return "", fmt.Errorf("числа должны быть из одной системы счисления")
	case 2:
		num1 := parseRomanNum(nums[0])
		num2 := parseRomanNum(nums[1])
		res, err = RomanCalc(num1, num2, operator)
		if err != nil {
			return "", err
		}
	}
	return res, nil
}

func parseArabNum(numStr string) (int, error) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, fmt.Errorf("некорректное число (%s)", numStr)
	}
	if num < 1 || num > 10 {
		return 0, fmt.Errorf("число должно быть от 1 до 10")

	}
	return num, nil
}

func parseRomanNum(numStr string) int {
	var num int
	for indx, roman := range romans {
		if strings.Contains(numStr, roman) {
			num = indx + 1
		}
	}
	return num
}

func arabCalc(num1 int, num2 int, operator string) (string, error) {
	var res int
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	}
	return strconv.Itoa(res), nil
}

func RomanCalc(num1 int, num2 int, operator string) (string, error) {
	var res int
	switch operator {
	case "+":
		res = num1 + num2
	case "-":
		if num1 > num2 {
			res = num1 - num2
		} else {
			return "", fmt.Errorf("в римской системе счисления отсутсвует ноль и отрицательные числа")
		}
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	}
	return romans[res-1], nil
}
