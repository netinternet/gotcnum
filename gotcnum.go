package gotcnum

import "strconv"

func Tcnum(n string) bool {
	var even, odd, total, digits, digit10, digit11 int
	if len(n) != 11 {
		return false
	}
	for i := 0; i < 11; i++ {
		str2int, _ := strconv.Atoi(string(n[i]))
		digits++
		if digits <= 9 {
			if digits%2 == 0 {
				even = even + str2int
			} else {
				odd = odd + str2int
			}
			total = total + str2int
		} else if digits == 10 {
			total = total + str2int
			digit10 = ((odd * 7) - even) % 10
			if digit10 != str2int {
				return false
			}
		} else if digits == 11 {
			digit11 = total % 10
			if digit11 != str2int {
				return false
			}
		}
	}
	return true
}
