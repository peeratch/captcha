package captcha

import (
	"fmt"
)

var (
	mathOperator = map[int]string{
		0: "+",
		1: "-",
		2: "*",
	}

	numberAsText = map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}
)

const (
	Format string = "%v %s %v"

	LeftFormatOperator  = 0
	RightFormatOperator = 1
)

type Captcha struct {
	Right        int
	Left         int
	MathOperator int
	Format       int
}

func NewCaptcha(format, left, mathOperator, right int) string {
	c := Captcha{
		right,
		left,
		mathOperator,
		format,
	}
	return c.generateCaptcha()
}

func (c *Captcha) generateCaptcha() string {
	m := map[int]string{
		LeftFormatOperator:  fmt.Sprintf(Format, c.Left, mathOperator[c.MathOperator], numberAsText[c.Right]),
		RightFormatOperator: fmt.Sprintf(Format, numberAsText[c.Left], mathOperator[c.MathOperator], c.Right),
	}
	return m[c.Format]
}
