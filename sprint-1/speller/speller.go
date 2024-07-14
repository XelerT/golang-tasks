//go:build !solution

package speller

const (
	hundred  = 100 
	thousand = 10    * hundred
	million  = 1_000 * thousand
	billion  = 1_000 * million
)

func spell_number(n uint8) string {
	number_spelling := map[uint8] string {
		1  : "one",
		2  : "two",
		3  : "three",
		4  : "four",
		5  : "five",
		6  : "six",
		7  : "seven",
		8  : "eight",
		9  : "nine",
	}

	return number_spelling[n]
}

func spell_special_nums(n uint8) string {
	ten_teens_spelling := map[uint8] string {
		10 : "ten",
		11 : "eleven",
		12 : "twelve",
		13 : "thirteen",
		14 : "fourteen",
		15 : "fifteen",
		16 : "sixteen",
		17 : "seventeen",
		18 : "eighteen",
		19 : "nineteen",
	}

	v, k := ten_teens_spelling[n] 
	if !k {
		return spell_number(n)
	}
	return v 
}

func spell20_99(n uint8) string {
	tens_spelling := map[uint8] string {
		2  : "twenty",
		3  : "thirty",
		4  : "forty",
		5  : "fifty",
		6  : "sixty",
		7  : "seventy",
		8  : "eighty",
		9  : "ninety",
	}
	if n % 10 != 0 {
		return tens_spelling[n / 10] + "-" + spell_number(n % 10)
	}
	return tens_spelling[n / 10]
}

func required_space(n uint8) string {
	if n == 0 {
		return ""
	}
	return " "
}

func spell_part(n uint64, spelling string, div uint64) string {
	res := spell(n / div)
	remain := n % div
	res += spelling + required_space(uint8(remain)) + spell(remain)

	return res
}

func spell(n uint64) string {
	var res string

	switch {
	case n / billion > 0:
		res += spell_part(n, " billion", billion)
	case n / million > 0:
		res += spell_part(n, " million", million)
	case n / thousand > 0:
		res += spell_part(n, " thousand", thousand)
	case n / hundred > 0:
		res += spell_part(n, " hundred", hundred)
	case n > 19 && n < 100:
		res += spell20_99(uint8(n))
	default:
		res += spell_special_nums(uint8(n))
	}

	return res
}

func Spell(n int64) string {
	var res string
	
	if n < 0 {
		res = "minus "
		n *= -1
	} else if n == 0 {
		return "zero"
	}
	return res + spell(uint64(n)) 
}
