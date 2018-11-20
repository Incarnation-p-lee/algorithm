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
	assert "github.com/attic-labs/noms/go/d"
)

func maxCommonDivisor(m, n uint) uint {
	assert.PanicIfTrue(m < 2 || n < 2)

	if m < n {
		m, n = n, m
	}

	for r := uint(1); r != 0; m, n = n, r {
		r = m % n
	}

	return m
}

// MaxCommonDivisor compute the max common divisor from given 2 integer
func MaxCommonDivisor(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	} else {
		return int(maxCommonDivisor(uint(m), uint(n)))
	}
}

// MaxCommonDivisorUnsigned compute the max common divisor from given 2 unsigned integer
func MaxCommonDivisorUnsigned(m, n uint) uint {
	if m == 0 || n == 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	} else {
		return maxCommonDivisor(m, n)
	}
}

// MaxCommonDivisorAsync will wait the value from channel, take two int and compute
// max common divisor. After finished, will send the result to given channel
func MaxCommonDivisorAsync(c chan int) {
	c <- MaxCommonDivisor(<-c, <-c)
}

// MaxCommonDivisorUnsignedAsync will wait the value from channel, take two uint
// and compute max common divisor. After finished, will send the result to given
// channel
func MaxCommonDivisorUnsignedAsync(c chan uint) {
	c <- MaxCommonDivisorUnsigned(<-c, <-c)
}
