package captcha

import (
	"captcha/strcon"
	"errors"
	"fmt"
)

var (
	ErrUnsupportMathOperator = errors.New("unsupport math operator")
	ErrUnsupportOperator     = errors.New("unsupport operator")

	mathOperator = map[string]string{
		"0": "+",
		"1": "-",
		"2": "*",
	}
)

const (
	Format string = "%s %s %s"

	LeftOperator  string = "0"
	RightOperator string = "1"
)

type Captcha struct {
	Right        string
	Left         string
	MathOperator string
	Operator     string
}

func NewCaptcha(right, left, mathOperator, operator string) *Captcha {
	return &Captcha{
		right,
		left,
		mathOperator,
		operator,
	}
}

func (c *Captcha) Validate() error {
	if _, ok := mathOperator[c.MathOperator]; !ok {
		return ErrUnsupportMathOperator
	}
	if c.Operator != LeftOperator && c.Operator != RightOperator {
		return ErrUnsupportOperator
	}
	return nil
}

func (c *Captcha) GenerateCaptcha() string {
	if c.Operator == LeftOperator {
		return c.doLeft()
	}
	return c.doRight()
}

func (c *Captcha) mathOperatorSymbol() string {
	return mathOperator[c.MathOperator]
}

func (c *Captcha) doLeft() string {
	return fmt.Sprintf(Format, c.Left, c.mathOperatorSymbol(), strcon.Ntot(c.Right))
}

func (c *Captcha) doRight() string {
	return fmt.Sprintf(Format, strcon.Ntot(c.Left), c.mathOperatorSymbol(), c.Right)
}
