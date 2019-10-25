package strcon

import (
	"strconv"

	"github.com/divan/num2words"
)

func Ntot(number string) string {
	return num2words.Convert(atoi(number))
}

func atoi(number string) int {
	i, _ := strconv.Atoi(number)
	return i
}
