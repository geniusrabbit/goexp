//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017
// @project GeniusRabbit 2017
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

func TestCondition(t *testing.T) {
	// (v1 = 'val1' OR v2 = 100) AND (v3 not in ('v1', 'v2', 'v3') OR v4 > 100)
	var q = And(
		Or(
			Func(func(val interface{}) bool {
				return val.(filter).v1 == "val1"
			}),
			Func(func(val interface{}) bool {
				return val.(filter).v2 == 100
			}),
		),
		Or(
			Not(Func(func(val interface{}) bool {
				var arr = []string{"v1", "v2", "v3"}
				for _, v := range arr {
					if v == val.(filter).v3 {
						return true
					}
				}
				return false
			})),
			Func(func(val interface{}) bool {
				return val.(filter).v4 > 100
			}),
		),
	)

	if !q.True(filter{v1: "val1", v2: 111, v3: "v4", v4: 1000}) {
		t.Error("Condition 1 error")
	}

	if !q.True(filter{v1: "val2", v2: 100, v3: "v4", v4: 100}) {
		t.Error("Condition 2 error")
	}

	if q.True(filter{v1: "val2", v2: 101, v3: "v3", v4: 1000}) {
		t.Error("Condition 3 error")
	}
}
