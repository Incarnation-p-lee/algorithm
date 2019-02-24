// MIT License
//
// Copyright (c) 2018 Pan Li
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package assert

import (
	"reflect"
	"testing"
)

type That struct {
	Value interface{}
	T     *testing.T
}

type Assertion interface {
	IsEqualsTo(expect interface{})
}

// IsEqualsTo will compare the expect value with type That, will trigger
// error if not equals.
func (d That) IsEqualsTo(expect interface{}) {
	switch expect := expect.(type) {
	case bool:

	case int:
	case int32:
	case int64:

	case uint:
	case uint32:
	case uint64:

	case float32:
	case float64:
		isEqualsToPrimitive(expect, &d)
	default:
		panic("Not implemented yet.")
	}
}

func isEqualsToPrimitive(expect interface{}, d *That) {
	ok := reflect.TypeOf(d.Value) == reflect.TypeOf(expect)
	v := d.Value

	if !ok {
		d.T.Errorf("Expecting instance type %T, but actually got %T", expect, v)
	} else if v != expect {
		d.T.Errorf("Expecting instance value %v, but actually got %v", expect, v)
	}
}
