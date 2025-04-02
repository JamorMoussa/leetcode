package main

// https://leetcode.com/problems/integer-to-roman/description/

// func decompose(num int) (int, int) {

// 	var div, carry int

// 	if num == 4 {
// 		div = 0
// 		carry = -1
// 	} else if num == 9 {
// 		div = 1
// 		carry = -1
// 	} else {
// 		div = num / 5
// 		carry = num % 5
// 	}

// 	return div, carry
// }

// func intToRoman(num int) string {

// 	roman := ""

// 	Exp := 1000

// 	intToRomanMap := map[int]string{
// 		1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M",
// 	}

// 	for num != 0 || Exp != 0 {

// 		div := num / Exp
// 		num %= Exp

// 		base, carry := decompose(div)

// 		if carry == -1 {

// 			first, second := "", ""
// 			first = intToRomanMap[Exp]

// 			if base == 0 {
// 				second = intToRomanMap[Exp*5]
// 			} else {
// 				second = intToRomanMap[Exp*10]
// 			}

// 			roman = roman + first + second
// 		}

// 		Exp /= 10
// 	}
// 	return roman
// }

type Tuple struct {
	First  string
	Second int
}

func intToRoman(num int) string {

	lookup := []Tuple{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	roman := ""

	for _, tuple := range lookup {

		for num >= tuple.Second {
			roman += tuple.First
			num -= tuple.Second
		}
	}

	return roman
}

// func main() {
// 	num := 3749

// 	fmt.Println(intToRoman(num))
// }
