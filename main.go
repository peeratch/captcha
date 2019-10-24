package main

import (
	"captcha/captcha"
	"fmt"
	"log"
)

// captcha ง่ายๆคือ รับ parameter 4 ตัว
// format ,left-operand, operator, right-operand

// operator มี 0, 1, 2 แทน +, - และคูณ

// format มี 2 แบบ
// ถ้าใส่ 0 เข้ามา จะแทน left-operand เป็นตัวเลข ส่วน right-operand เป็นตัวหนังสือเช่น
// 0, 1, 1, 1 จะได้ “1 + one”
// 1, 1, 1, 1 จะได “one + 1"

func main() {
	c := captcha.NewCaptcha("1", "1", "0", "1")
	if err := c.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(c.Process())
}
