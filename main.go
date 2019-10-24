package main

import (
	"captcha/captcha"
	"fmt"
	"log"
)

func main() {
	c := captcha.NewCaptcha("1", "1", "0", "1")
	if err := c.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Print(c.GenerateCaptcha())
}
