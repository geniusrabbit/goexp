//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017, 2023
// @project GeniusRabbit 2017, 2023
// @license Apache-2.0
//

package goexp

// Not Condition implementation
type not[T any] struct {
	C Condition[T]
}

// Not creates new not condition
func Not[T any](c Condition[T]) Condition[T] {
	return not[T]{C: c}
}

// True if Condition executed
func (n not[T]) True(val T) bool {
	return !n.C.True(val)
}

// And logic Condition implementation
type and[T any] struct {
	C []Condition[T]
}

// And creates new AND condition
func And[T any](c ...Condition[T]) and[T] {
	return and[T]{C: c}
}

// True if Condition executed
func (a and[T]) True(val T) bool {
	for _, c := range a.C {
		if !c.True(val) {
			return false
		}
	}
	return true
}

// Or logic Condition implementation
type or[T any] struct {
	C []Condition[T]
}

// Or creates new OR condition
func Or[T any](c ...Condition[T]) Condition[T] {
	return or[T]{C: c}
}

// True if Condition executed
func (o or[T]) True(val T) bool {
	for _, c := range o.C {
		if c.True(val) {
			return true
		}
	}
	return false
}

type tfunc[T any] struct {
	f func(T) bool
}

// Func creates condition procedure
func Func[T any](f func(T) bool) Condition[T] {
	return tfunc[T]{f: f}
}

// True if Condition executed
func (f tfunc[T]) True(val T) bool {
	return f.f(val)
}

type extr[T any, E any] struct {
	extract func(T) E
	cond    Condition[E]
}

// Extr defines context extract operation
func Extr[T any, E any](extract func(T) E, cond Condition[E]) Condition[T] {
	return extr[T, E]{extract: extract, cond: cond}
}

// True if Condition executed
func (e extr[T, C]) True(val T) bool {
	return e.cond.True(e.extract(val))
}
