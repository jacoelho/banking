package bic_test

import (
	"fmt"
	"github.com/jacoelho/banking/bic"
)

func ExampleBIC_IsValid() {
	// Parse a BIC
	b, err := bic.Parse("ABCDBEB1XXX")
	fmt.Println(err == nil && b.IsValid())
	// Output: true
}
