package main

import (
	"fmt"
	"strconv"
)

//输入: num1 = "123", num2 = "456"
//输出: "56088"

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))

}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	result := "0"
	for i := l1 - 1; i >= 0; i-- {
		tmp := make([]string, 0)
		for n := 0; n < l1-1-i; n++ {
			tmp = append(tmp, "0")
		}
		extra := 0
		for j := l2 - 1; j >= 0; j-- {
			current := (int(num1[i])-'0')*(int(num2[j])-'0') + extra
			tmp = append(tmp, strconv.Itoa(current%10))
			extra = current / 10
		}
		if extra > 0 {
			tmp = append(tmp, strconv.Itoa(extra))
		}
		tmpRet := ""
		for l := len(tmp); l > 0; l-- {
			tmpRet += tmp[l-1]
		}
		result = add(result, tmpRet)
	}

	return result
}

func add(num1 string, num2 string) string {
	if num1 == "0" {
		return num2
	}
	if num2 == "0" {
		return num1
	}

	l1 := len(num1)
	l2 := len(num2)
	extra := 0
	ret := make([]string, 0, len(num1))
	for i := 0; i < l1 || i < l2; i++ {
		a := 0
		b := 0
		if l1 > i {
			a = int(num1[l1-i-1] - '0')
		}
		if l2 > i {
			b = int(num2[l2-i-1] - '0')
		}
		sum := a + b + extra
		extra = 0
		if sum >= 10 {
			extra = 1
		}
		ret = append(ret, strconv.Itoa(sum%10))
	}
	if extra > 0 {
		ret = append(ret, "1")
	}
	result := ""
	for l := len(ret); l > 0; l-- {
		result += ret[l-1]
	}
	return result
}
