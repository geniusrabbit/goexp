//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017, 2023
// @project GeniusRabbit 2017, 2023
// @license Apache-2.0
//

package goexp

import "testing"

type filter struct {
	v1 string
	v2 int
	v3 string
	v4 int
}

type filterV12 struct {
	v1 string
	v2 int
}

// (v1 = 'val1' OR v2 = 100) AND (v3 not in ('v1', 'v2', 'v3') OR v4 > 100)
var testCondition = And(
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

var tests = []struct {
	c *filter
	r bool
}{
	{c: &filter{v1: "val1", v2: 111, v3: "v4", v4: 1000}, r: true},
	{c: &filter{v1: "val2", v2: 100, v3: "v4", v4: 100}, r: true},
	{c: &filter{v1: "val2", v2: 101, v3: "v3", v4: 1000}, r: false},
}

func TestCondition(t *testing.T) {
	for i, varCtx := range tests {
		if testCondition.True(varCtx.c) != varCtx.r {
			t.Errorf("Condition %d error", i+1)
		}
	}
}

func BenchmarkCondition(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		cnt := len(tests)
		for p.Next() {
			testCondition.True(tests[b.N%cnt].c)
		}
	})
}
