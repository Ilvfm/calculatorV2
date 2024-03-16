package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Приветствую вас,я будущий супер калькулятор\n" +
		"Правила моего использования:\n" +
		"1. Введите в меня два числа и знак между ними.\n" +
		"2. Пока что могу и арабские,и римские считать\n" +
		"3. Но учтите,только такие знаки воспринимаю: '+', '-', '*', '/'\n" +
		"\n\n" +
		"Введите пример:")
	intType, first, second, sign, err := readLine()
	if err != nil {
		fmt.Println("Ошибочка,очень жаль:\n", err)
		return
	}
	if intType == "arab" {
		firstNum, err1 := strconv.Atoi(first)
		if err1 != nil {
			fmt.Println("И вновь какая-то ошибка,что же делать:\n", err1)
			return
		}
		secondNum, err2 := strconv.Atoi(second)
		if err2 != nil {
			fmt.Println("Исправьте меня,пожалуйста:\n", err2)
			return
		}
		res, err3 := calculator(firstNum, secondNum, sign)
		if err3 != nil {
			fmt.Println("Я чуть-чуть перегрелся:\n", err3)
			return
		} else {
			fmt.Println("Ответ: ", res)
		}
	} else {
		firstNum := fromRomanToInt(first)
		secondNum := fromRomanToInt(second)
		res, err1 := calculator(firstNum, secondNum, sign)
		if err1 != nil {
			fmt.Println("Ошибочка:\n", err1)
			return
		} else {
			final, err2 := fromIntToRoman(res)
			if err2 != nil {
				fmt.Println("Ошибочка:\n", err2)
				return
			}
			fmt.Println("Ответ: ", final)
		}
	}
}
func makepPanic() {
	defer func() {
		panicValue := recover()
		fmt.Println(panicValue)
	}()
	panic("Я так больше не могу!!!")
}

func calculator(first int, second int, sign string) (int, error) {
	if first > 10 || second > 10 {
		return 8, errorHandler(8)
	}
	switch {
	case sign == "+":
		return first + second, nil
	case sign == "-":
		return first - second, nil
	case sign == "*":
		return first * second, nil
	case sign == "/" && second != 0:
		return first / second, nil
	case sign == "/" && second == 0:
		return 4, errorHandler(4)
	default:
		return 5, errorHandler(5)
	}
}
func readLine() (string, string, string, string, error) {
	stdin := bufio.NewReader(os.Stdin)
	usInput, _ := stdin.ReadString('\n')
	usInput = strings.TrimSpace(usInput)
	intType, first, second, sign, err := checkInput(usInput)
	if err != nil {
		return "", "", "", "", err
	}
	return intType, first, second, sign, err
}

func checkInput(input string) (string, string, string, string, error) {
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(input, "")
	arr := strings.Split(replace, "")
	var intType, first, second, sign string
	for index, value := range arr {
		isN := isNumber(value)
		isS := isSign(value)
		isR := isRomanNumber(value)
		if !isN && !isS && !isR {
			return "", "", "", "", errorHandler(1)
		}
		if isS {
			if sign != "" {
				return "", "", "", "", errorHandler(6)
			} else {
				sign = arr[index]
			}
		}
		if (isN && intType != "roman") || (isR && intType != "arab") {
			if intType == "" {
				if isN {
					intType = "arab"
				} else {
					intType = "roman"
				}
			}
			if first == "" && !(index+1 == len(arr)) && isSign(arr[index+1]) {
				slice := arr[0:(index + 1)]
				first = strings.Join(slice, "")
			} else if index+1 == len(arr) && first != "" {
				slice := arr[(len(first) + 1):]
				second = strings.Join(slice, "")
			}
		} else if (intType == "arab" && isR) || (intType == "roman" && isN) {
			return "", "", "", "", errorHandler(2)
		}
	}
	if second == "" || first == "" || sign == "" {
		return "", "", "", "", errorHandler(3)
	}
	return intType, first, second, sign, nil
}

func isNumber(c string) bool {
	if c >= "0" && c <= "9" {
		return true
	} else {
		return false
	}
}

func isSign(c string) bool {
	if c == "+" || c == "-" || c == "/" || c == "*" {
		return true
	} else {
		return false
	}
}
func isRomanNumber(c string) bool {
	_, ok := dict[c]
	if ok {
		return true
	} else {
		return false
	}
}

func errorHandler(code int) error {
	return errors.New(errorDict[code])
}

var errorDict = map[int]string{
	1: "Посмотрите правила ещё раз и используйте только эти операторы: '+', '-', '/', '*' ",
	2: "Что-то пошло не так ",
	3: "Я слабенький,могу только два числа и знак обработать",
	4: "Не умею делить на ноль, но научусь!!!",
	5: "Одну секунду",
	6: "Один пример,пожааалуйста",
	7: "А вот в римской нет отрицательных чисел",
	8: "Пожалуйста, введите числа от 0 до 10",
}

var dict = map[string]int{
	"M":  1000,
	"CM": 900,
	"D":  500,
	"CD": 400,
	"C":  100,
	"XC": 90,
	"L":  50,
	"XL": 40,
	"XX": 20,
	"X":  10,
	"IX": 9,
	"V":  5,
	"IV": 4,
	"I":  1,
}

func fromRomanToInt(roman string) int {
	var res int
	arr := strings.Split(roman, "")
	for index, value := range arr {
		if index+1 != len(arr) && dict[value] < dict[arr[index+1]] {
			res -= dict[value]
		} else {
			res += dict[value]
		}
	}
	return res
}

func fromIntToRoman(number int) (string, error) {
	if number <= 0 {
		return "", errorHandler(7)
	}
	arr1 := [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	arr2 := [13]string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var str string
	for number > 0 {
		for i := 0; i < 13; i++ {
			if arr1[i] <= number {
				str += arr2[i]
				number -= arr1[i]
				break
			}
		}
	}
	return str, nil
}
