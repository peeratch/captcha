package captcha

import (
	"errors"
	"testing"
)

func TestCaptcha_Validator(t *testing.T) {
	t.Run("captcha validate should return error not support operator", func(t *testing.T) {
		expecrted := errors.New("unsupport operator")
		c := Captcha{
			Right:        "1",
			Left:         "1",
			MathOperator: "0",
			Operator:     "3"}

		actual := c.Validate()
		if expecrted.Error() != actual.Error() {
			t.Errorf("expected error: %s but got %s", expecrted, actual)
		}
	})

	t.Run("captcha validate should return error not support math operator", func(t *testing.T) {
		expecrted := errors.New("unsupport math operator")
		c := Captcha{
			Right:        "1",
			Left:         "1",
			MathOperator: "4",
			Operator:     "1"}

		actual := c.Validate()
		if expecrted.Error() != actual.Error() {
			t.Errorf("expected error: %s but got %s", expecrted, actual)
		}
	})

	t.Run("captcha validate should return passed", func(t *testing.T) {
		c := Captcha{
			Right:        "1",
			Left:         "1",
			MathOperator: "0",
			Operator:     "1"}

		actual := c.Validate()
		if nil != actual {
			t.Errorf("expected nil but got %s", actual)
		}
	})
}

func TestCaptcha_ConvertString_To_MathSymbol(t *testing.T) {
	expected := []string{"+", "-", "*"}
	actual := []Captcha{
		Captcha{MathOperator: "0"},
		Captcha{MathOperator: "1"},
		Captcha{MathOperator: "2"},
	}
	for index := 0; index < len(actual); index++ {
		t.Run("convert string to math symbol shold return correct result", func(t *testing.T) {
			if expected[index] != actual[index].toMathOperatorSymbol() {
				t.Errorf("expected: %s should equal actual but got %s", expected[index], actual[index])
			}
		})
	}
}

func TestCaptcha_ShouldDoLeft_Operation(t *testing.T) {
	expected := "1 + One"
	c := NewCaptcha("1", "1", "0", "0")
	if err := c.Validate(); err != nil {
		t.Errorf("captcha process should not got error but got %s", err)
	}
	acutal := c.Process()
	if expected != acutal {
		t.Errorf("captcha process should return %s but got %s", expected, acutal)
	}

}

func TestCaptcha_ShouldDoRight_Operation(t *testing.T) {
	expected := "One + 1"
	c := NewCaptcha("1", "1", "0", "1")
	if err := c.Validate(); err != nil {
		t.Errorf("captcha process should not got error but got %s", err)
	}
	acutal := c.Process()
	if expected != acutal {
		t.Errorf("captcha process should return %s but got %s", expected, acutal)
	}
}
