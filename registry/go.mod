module github.com/jacoelho/banking/registry

go 1.24.0

replace github.com/jacoelho/banking => ../

require (
	github.com/google/go-cmp v0.4.0
	github.com/jacoelho/banking v0.0.0-00010101000000-000000000000
	gopkg.in/yaml.v2 v2.4.0
)

require golang.org/x/text v0.31.0
