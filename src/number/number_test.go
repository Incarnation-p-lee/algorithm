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

package number

import (
	"math"
	"testing"
)

func TestMaxCommonDivisor(t *testing.T) {
	var data = []struct {
		a, b   int
		expect int
	}{
		{-1, 0, 0},
		{-1, -1, 0},
		{-8, 1, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
		{2, 3, 1},
		{119, 342, 1},
		{111, 999, 111},
		{18, 15, 3},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		{2147483647, 30, 1},
		{math.MaxInt32, 1024, 1},
		{1001, 100101, 1},
		{24, 18, 6},
		{49, 63, 7},
	}

	for _, d := range data {
		actual := MaxCommonDivisor(d.a, d.b)

		if actual != d.expect {
			t.Errorf("MaxCommonDivisor(%v, %v) == %v, but actual %v.",
				d.a, d.b, d.expect, actual)
		}
	}
}

func TestMaxCommonDivisorUnsigned(t *testing.T) {
	var data = []struct {
		a, b   uint
		expect uint
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{1, 0, 0},
		{0, 1, 0},
		{2, 3, 1},
		{119, 342, 1},
		{111, 999, 111},
		{18, 15, 3},
		{math.MaxUint32, math.MaxUint32, math.MaxUint32},
		{2147483647, 30, 1},
		{math.MaxUint32, 1024, 1},
		{1001, 100101, 1},
		{24, 18, 6},
		{49, 63, 7},
	}

	for _, d := range data {
		actual := MaxCommonDivisorUnsigned(d.a, d.b)

		if actual != d.expect {
			t.Errorf("MaxCommonDivisorUnsigned(%v, %v) == %v, but actual %v.",
				d.a, d.b, d.expect, actual)
		}
	}
}

func TestMaxCommonDivisorAsync(t *testing.T) {
	var data = []struct {
		a, b   int
		expect int
	}{
		{-1, 0, 0},
		{-1, -1, 0},
		{-8, 1, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
		{2, 3, 1},
		{119, 342, 1},
		{111, 999, 111},
		{18, 15, 3},
		{math.MaxInt32, math.MaxInt32, math.MaxInt32},
		{2147483647, 30, 1},
		{math.MaxInt32, 1024, 1},
		{1001, 100101, 1},
		{24, 18, 6},
		{49, 63, 7},
	}

	c := make(chan int)
	defer close(c)

	for _, d := range data {
		go MaxCommonDivisorAsync(c)

		c <- d.a
		c <- d.b
		actual := <-c

		if actual != d.expect {
			t.Errorf("MaxCommonDivisorAsync(%v, %v) == %v, but actual %v.",
				d.a, d.b, d.expect, actual)
		}
	}
}
