package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var operation string
var result string

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Println("Input:")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		WorkingString(text)
	}

}
func Operation(s string) string {
	var op [40]string
	for i := 0; i < len(s); i++ {
		op[i] = string(s[i])
		if op[i] == "+" {
			operation = "+"
		} else if op[i] == "-" {
			operation = "-"
		} else if op[i] == "*" {
			operation = "*"
		} else if op[i] == "/" {
			operation = "/"
		}
	}
	return operation
}
func Calculator(num []int, str []string, op string) {
	switch op {
	case "+":
		if len(str) == 2 && len(num) == 0 {
			result = str[0] + str[1]
		} else {
			panic("Необходимо складывать строки")
		}

	case "-":
		if len(str) == 2 && len(num) == 0 {

			result = strings.Join(strings.SplitN(str[0], str[1], 2), "")
		} else {
			panic("Необходимо вычитать строки")
		}

	case "*":
		if len(str) == 1 && len(num) == 1 {
			for i := 0; i < num[0]; i++ {
				result = str[0] + result
			}
		} else {

			panic("Необходимо уможать строку на число")
		}
	case "/":
		if len(str) == 1 && len(num) == 1 {
			result = str[0][:len(str[0])/num[0]]
		} else {
			panic("Необходимо делить строку на число")
		}
	}
	if len(result) > 40 {
		fmt.Println("Output:" + "\"" + result[0:40] + "\"" + "...")

	} else {
		fmt.Println("Output:" + "\"" + result + "\"")
	}
}
func WorkingString(s string) {
	numbers := make([]int, 0)
	stringS := make([]string, 0)
	var str = regexp.MustCompile("[+-/*]").Split(s, -1)
	if _, err := strconv.Atoi(str[0]); err == nil {
		panic("Первым значением должно быть слово")
	}
	for _, elem := range str {
		num, err := strconv.Atoi(elem)
		if err != nil {
			if (elem[0] & elem[len(elem)-1]) != '"' {
				panic("Вводить данные без пробелов между знаков и слова должны быть в ковычках")
			}
			elem = elem[1 : len(elem)-1] //убираем ковычки
			if len(elem) > 30 {
				panic("Слова должны быть не длиннее 10 символов")
			}
			stringS = append(stringS, elem)
		} else {
			numbers = append(numbers, num)
		}
	}
	if !((len(stringS) == 2) && (len(numbers) == 0)) && !((len(stringS) == 1) && (len(numbers) == 1)) {
		panic("Надо воодить не больше 2 переменных")
	}
	Calculator(numbers, stringS, Operation(s))
}
