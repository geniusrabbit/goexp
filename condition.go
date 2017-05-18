//
// @author Dmitry Ponomarev <demdxx@gmail.com> 2017
// @project GeniusRabbit 2017
// @license Apache-2.0
//

package goexp

// Condition interface declaration
type Condition interface {
	True(val interface{}) bool
}
