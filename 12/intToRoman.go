package main

import "fmt"

func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return string(num)
	}
	result := ""
	tmp := 0
	for i := 1000; i > 0; i = i / 10 {
		tmp, num = num/i, num%i
		fmt.Println(tmp, num)
		if tmp <= 0 {
			continue
		}
		switch i {
		case 1000:
			for tmp > 0 {
				result += "M"
				tmp--
			}
		case 100:
			result += count(tmp,"M", "D", "C")
		case 10:
			result += count(tmp,"C", "L", "X")
		case 1:
			result += count(tmp,"X", "V", "I")
		default:
		}
	}
	return result
}

func count(tmp int,ten, five, one string) (result string) {
	if tmp == 9 {
		result += one + ten
	} else if tmp >= 5 {
		j := five
		for tmp > 5 {
			j += one
			tmp--
		}
		result += j
	} else if tmp == 4 {
		result += one + five
	} else {
		for tmp > 0 {
			result += one
			tmp--
		}
	}
	return result
}
func main() {
	fmt.Println(intToRoman(1994))
}
