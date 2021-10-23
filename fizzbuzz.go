package fizzbuzz

import "fmt"

func fizzBuzz(n int) []string {
	result := make([]string, 0)
	for i := 1; i <= n; i++ {
		var v string
		if (i % 3 == 0) && (i % 5 == 0) {
			v = "fizzbuzz"
		} else if i % 3 == 0 {
			v = "fizz"
		} else if i % 5 == 0 {
			v = "buzz"
		} else {
			v = fmt.Sprint(i)
		}
		result = append(result, v)
	}
	return result
}
