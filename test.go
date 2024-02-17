package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	dict := make(map[string]int)

	dict["I"] = 1
	dict["II"] = 2
	dict["III"] = 3
	dict["IV"] = 4
	dict["V"] = 5
	dict["VI"] = 6
	dict["VII"] = 7
	dict["VIII"] = 8
	dict["IX"] = 9
	dict["X"] = 10

	fmt.Print("Введите действие: ")

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name := scanner.Text()
	var reshenie = do_everything(name, dict)

	fmt.Println(name, "=", reshenie)
}

func do_everything(str string, dict map[string]int) string {
	var resh int

	plus := strings.Split(str, "+")
	if len(plus) == 2 {
		if dict[plus[0]] != 0 && dict[plus[1]] != 0 {
			return with_rome(plus[0], plus[1], "+", dict)
		} else {
			num1, err := strconv.Atoi(plus[0])
			num2, err := strconv.Atoi(plus[1])
			if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
				panic("wrong num")
			}
			resh = num1 + num2
			if err != nil {
				panic(err)
			}
		}
		return strconv.Itoa(resh)
	} else {

		minus := strings.Split(str, "-")
		if len(minus) == 2 {
			if dict[minus[0]] != 0 && dict[minus[1]] != 0 {
				return with_rome(minus[0], minus[1], "-", dict)
			} else {
				num1, err := strconv.Atoi(minus[0])
				num2, err := strconv.Atoi(minus[1])
				if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
					panic("wrong num")
				}
				resh = num1 + (-num2)
				if err != nil {
					panic(err)
				}
			}
			return strconv.Itoa(resh)
		} else {

			umnoshit := strings.Split(str, "*")
			if len(umnoshit) == 2 {
				if dict[umnoshit[0]] != 0 && dict[umnoshit[1]] != 0 {
					return with_rome(umnoshit[0], umnoshit[1], "*", dict)
				} else {

					num1, err := strconv.Atoi(umnoshit[0])
					num2, err := strconv.Atoi(umnoshit[1])
					if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
						panic("wrong num")
					}
					resh = num1 * num2
					if err != nil {
						panic(err)
					}
				}
				return strconv.Itoa(resh)
			} else {

				delit := strings.Split(str, "/")
				if len(delit) == 2 {
					if dict[delit[0]] != 0 && dict[delit[1]] != 0 {
						return with_rome(delit[0], delit[1], "/", dict)
					} else {
						num1, err := strconv.Atoi(delit[0])
						num2, err := strconv.Atoi(delit[1])
						if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
							panic("wrong num")
						}
						resh = num1 / num2
						if err != nil {
							panic(err)
						}
					}
					return strconv.Itoa(resh)
				} else {
					panic("i cant find +-*/ or many nums")
				}
			}
		}
	}
}

func with_rome(first, second, znak string, dict map[string]int) string {
	frst := dict[first]
	scnd := dict[second]
	var reshenie int
	if znak == "+" {
		reshenie = frst + scnd
	} else if znak == "-" {
		reshenie = frst - scnd
	} else if znak == "*" {
		reshenie = frst * scnd
	} else {
		reshenie = frst / scnd
	}

	var back string

	if reshenie <= 0 {
		panic("<0")
	} else if reshenie <= 10 {
		return find_key(reshenie, dict)
	} else if reshenie <= 20 {
		back = "X"
		reshenie -= 10
		back += find_key(reshenie, dict)
		return back
	} else if reshenie <= 30 {
		back = "XX"
		reshenie -= 20
		back += find_key(reshenie, dict)
		return back
	} else if reshenie < 40 {
		back = "XXX"
		reshenie -= 30
		back += find_key(reshenie, dict)
		return back
	} else if reshenie == 40 {
		return "XL"
	} else if reshenie < 50 {
		back = "XL"
		reshenie -= 40
		back += find_key(reshenie, dict)
		return back
	} else if reshenie == 50 {
		return "L"
	} else if reshenie <= 60 {
		back = "L"
		reshenie -= 50
		back += find_key(reshenie, dict)
		return back
	} else if reshenie <= 70 {
		back = "LX"
		reshenie -= 60
		back += find_key(reshenie, dict)
		return back
	} else if reshenie <= 80 {
		back = "LXX"
		reshenie -= 70
		back += find_key(reshenie, dict)
		return back
	} else if reshenie < 90 {
		back = "LXXX"
		reshenie -= 80
		back += find_key(reshenie, dict)
		return back
	} else if reshenie == 90 {
		return "XC"
	} else if reshenie < 100 {
		back = "XC"
		reshenie -= 90
		back += find_key(reshenie, dict)
		return back
	} else {
		return "C"
	}
}

func find_key(arg int, dict map[string]int) string {
	for key, val := range dict {
		if val == arg {
			return key
		}
	}
	panic("strange num")
}
