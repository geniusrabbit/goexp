# goexp

![License](https://img.shields.io/github/license/geniusrabbit/goexp)
[![Go Report Card](https://goreportcard.com/badge/github.com/geniusrabbit/goexp)](https://goreportcard.com/report/github.com/geniusrabbit/goexp)
[![Coverage Status](https://coveralls.io/repos/github/geniusrabbit/goexp/badge.svg?branch=main)](https://coveralls.io/github/geniusrabbit/goexp?branch=main)
[![Testing Status](https://github.com/geniusrabbit/goexp/workflows/Tests/badge.svg)](https://github.com/geniusrabbit/goexp/actions?workflow=Tests)
[![Publish Docker Status](https://github.com/geniusrabbit/goexp/workflows/Publish/badge.svg)](https://github.com/geniusrabbit/goexp/actions?workflow=Publish)

The simple builder structures of expression

## Condition example

> (v1 = 'val1' OR v2 = 100) AND (v3 not in ('v1', 'v2', 'v3') OR v4 > 100)
```go
condition1 := And(
	Extr(
		func(val *filter) *filterV12 {
			return &filterV12{v1: val.v1, v2: val.v2}
		},
		Or(
			Func(func(val *filterV12) bool {
				return val.v1 == "val1"
			}),
			Func(func(val *filterV12) bool {
				return val.v2 == 100
			}),
		),
	),
	Or(
		Not(Func(func(val *filter) bool {
			var arr = []string{"v1", "v2", "v3"}
			for _, v := range arr {
				if v == val.v3 {
					return true
				}
			}
			return false
		})),
		Func(func(val *filter) bool {
			return val.v4 > 100
		}),
	),
)

if condition1.True(&filter{...}) {
  // DO SOMETHING ...
}
```

> Combination of condition expressions
```go
conditionNew := Extr(
  func(ctx any) *filter {
    return &filter{...}
  },
  And(
    condition1,
    Or(
      condition2,
      condition3,
    ),
  ),
)
```
