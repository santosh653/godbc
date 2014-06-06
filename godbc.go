//
// Copyright 2014 Luis Pabon, Jr.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Design-by-Contract for Go
//
// Design by Contract is a programming methodology
// which binds the caller and the function called to a
// contract. The contract is represented using Hoare Triple:
//      {P} C {Q}
// where {P} is the precondition before executing command C,
// and {Q} is the postcondition.
//
// See Also
//
// * http://en.wikipedia.org/wiki/Design_by_contract
// * http://en.wikipedia.org/wiki/Hoare_logic
// * http://dlang.org/dbc.html
//
package godbc

import (
	"errors"
	"fmt"
	"runtime"
)

// InvariantSimpleTester is an interface which provides a receiver to
// test the object
type InvariantSimpleTester interface {
	Invariant() bool
}

// InvariantTester is an interface which provides not only an Invariant(),
// but also a receiver to print the structure
type InvariantTester interface {
	InvariantSimpleTester
	String() string
}

// dbc_panic prints to the screen information of the failure followed
// by a call to panic()
func dbc_panic(dbc_func_name string, b bool, message ...interface{}) {
	if !b {

		// Get caller information which is the caller
		// of the caller of this function
		pc, file, line, _ := runtime.Caller(2)
		caller_func_info := runtime.FuncForPC(pc)

		error_string := fmt.Sprintf("%s:\n\r\tfunc (%s) 0x%x\n\r\tFile %s:%d",
			dbc_func_name,
			caller_func_info.Name(),
			pc,
			file,
			line)

		if len(message) > 0 {
			error_string += fmt.Sprintf("\n\r\tInfo: %+v", message)
		}
		err := errors.New(error_string)

		// Finally panic
		panic(err)
	}
}

// Require checks that the preconditions are satisfied before
// executing the function
//
// Example
//
// 		func Divide(a, b int) int {
//			godbc.Require(b != 0)
//			return a/b
// 		}
//
func Require(b bool, message ...interface{}) {
	dbc_panic("REQUIRE", b, message...)
}

// Ensure checks the postconditions are satisfied before returning
// to the caller.
//
// Example Code
//
//		type Data struct {
//			a int
//		}
//
// 		func (*d Data) Set(a int) {
//			d.a = a
//			godbc.Ensure(d.a == a)
// 		}
//
func Ensure(b bool, message ...interface{}) {
	dbc_panic("ENSURE", b, message...)
}

// Check provides a simple assert
func Check(b bool, message ...interface{}) {
	dbc_panic("CHECK", b, message...)
}

// InvariantSimple calls the objects Invariant() receiver to test
// the object for correctness.
//
// The caller object must provide an object that supports the
// interface InvariantSimpleTester and does not need to provide
// a String() receiver
func InvariantSimple(obj InvariantSimpleTester, message ...interface{}) {
	dbc_panic("INVARIANT", obj.Invariant(), message...)
}

// Invariant calls the objects Invariant() receiver to test
// the object for correctness.
//
// The caller object must provide an object that supports the
// interface InvariantTester
//
// To see an example, please take a look at the godbc_test.go

func (d *Date) Invariant() bool {
	if (1 <= d.day && d.day <= 31) &&
		(1 <= d.month && d.month <= 12) {
		return true
	}
	return false
}

func (d *Date) Set(day, month int) {
	d.day, d.month = day, month
}

func (d *Date) String() string {
	return fmt.Sprintf("Day:%d Month:%d",
		d.day, d.month)
}

func TestInvariant(t *testing.T) {
	d := &Date{0, 0}
	assert.Panics(t, func() {
		Invariant(d)
	})

func Invariant(obj InvariantTester, message ...interface{}) {
	m := append(message, obj)
	dbc_panic("INVARIANT", obj.Invariant(), m)
}
