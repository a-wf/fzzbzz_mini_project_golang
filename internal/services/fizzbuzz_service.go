package services

import (
	"strconv"
	"strings"
)


func FizzBuzzAnswer (int1 int, int2 int, limit int, str1 string, str2 string) string{
		var fizzbuzz []string;
		
		for i := 1;  i<=limit; i++ {
			if i % int1 == 0 && i % int2 == 0 {
				fizzbuzz= append(fizzbuzz, str1 + str2)
			}	else if i % int1 == 0 {
				fizzbuzz= append(fizzbuzz, str1)

			} else if i % int2 == 0 {
				fizzbuzz= append(fizzbuzz, str2)
			} else {
				s := strconv.Itoa(i)
				fizzbuzz= append(fizzbuzz, s)
			}
		}

	return strings.Join(fizzbuzz, ",")	;
}