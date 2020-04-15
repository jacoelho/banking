package iban

import (
	"math/rand"

	"github.com/jacoelho/banking/ascii"
)

var generator = ascii.NewGenerator(nil)

func SeedGenerator(r *rand.Rand) {
	generator = ascii.NewGenerator(r)
}
