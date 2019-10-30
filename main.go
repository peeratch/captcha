package main

import (
	"captcha/captcha"
	"fmt"
)

func main() {
	fmt.Print(captcha.NewCaptcha(0, 1, 1, 1))
}
