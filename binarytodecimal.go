package proj

import "math"

func BinaryToDecimal(num int64) int64 {
	var remainder int64
	var index int64 = 0
	var decimalNum int64 = 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int64(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
