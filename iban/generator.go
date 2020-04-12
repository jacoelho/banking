package iban

import "github.com/jacoelho/banking/ascii"

var generator *ascii.Generator

func init() {
	generator = ascii.New(nil)
}
