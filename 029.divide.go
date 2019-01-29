package leetcode

const (
	INT_MIN = -2147483648
	INT_MAX = 2147483647
)

func divide(dividend int, divisor int) int {
	isDiff := false
	res := 0
	if dividend < 0 {
		dividend = -dividend
		isDiff = !isDiff
	}
	if divisor < 0 {
		divisor = -divisor
		isDiff = !isDiff
	}
	for dividend >= divisor {
		tmp := divisor
		d := 1
		for dividend >= (tmp << 1) {
			tmp = tmp << 1
			d = d << 1
		}
		dividend -= tmp
		res += d
	}
	if isDiff {
		res = -res
	}
	if res > INT_MAX || res < INT_MIN {
		res = INT_MAX
	}
	return res
}
