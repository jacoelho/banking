package rule

type Rule interface {
	Type() string
	String() string
}
