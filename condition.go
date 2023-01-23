//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017, 2023
// @project GeniusRabbit 2017, 2023
// @license Apache-2.0
//

package goexp

// Condition interface declaration
type Condition[T any] interface {
	True(val T) bool
}
