package ascii

import (
	"math/rand"
	"strings"
	"time"
)

const (
	digits           = "0123456789"
	lowerCaseLetters = "abcdefghijklmnopqrstuvwxyz"
	upperCaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphaNumeric     = digits + lowerCaseLetters + upperCaseLetters
)

type Generator struct {
	rand *rand.Rand
}

func New(r *rand.Rand) *Generator {
	if r == nil {
		r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}

	return &Generator{
		rand: r,
	}
}
func (g *Generator) Digits(sb *strings.Builder, length int) {
	g.stringWithCharset(sb, digits, length)
}

func (g *Generator) LowerCaseLetters(sb *strings.Builder, length int) {
	g.stringWithCharset(sb, lowerCaseLetters, length)
}

func (g *Generator) UpperCaseLetters(sb *strings.Builder, length int) {
	g.stringWithCharset(sb, upperCaseLetters, length)
}

func (g *Generator) AlphaNumeric(sb *strings.Builder, length int) {
	g.stringWithCharset(sb, alphaNumeric, length)
}

func (g *Generator) stringWithCharset(sb *strings.Builder, charset string, length int) {
	for i := 0; i < length; i++ {
		sb.WriteRune(rune(charset[g.rand.Intn(len(charset))]))
	}
}
