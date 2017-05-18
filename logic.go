//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017
// @project GeniusRabbit 2017
// @license Apache-2.0
//

package goexp

// Not Condition implementation
type not struct {
	C Condition
}

// Not creates new not condition
func Not(c Condition) Condition {
	return not{C: c}
}

// True if Condition executed
func (n not) True(val interface{}) bool {
	return !n.C.True(val)
}

// And logic Condition implementation
type and struct {
	C []Condition
}

// And creates new AND condition
func And(c ...Condition) Condition {
	return and{C: c}
}

// True if Condition executed
func (a and) True(val interface{}) bool {
	for _, c := range a.C {
		if !c.True(val) {
			return false
		}
	}
	return true
}

// Or logic Condition implementation
type or struct {
	C []Condition
}

// Or creates new OR condition
func Or(c ...Condition) Condition {
	return or{C: c}
}

// True if Condition executed
func (o or) True(val interface{}) bool {
	for _, c := range o.C {
		if c.True(val) {
			return true
		}
	}
	return false
}

// Func condition procedure
type Func func(val interface{}) bool

// True if Condition executed
func (f Func) True(val interface{}) bool {
	return f(val)
}
