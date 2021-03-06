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
	"testing"
)

func TestIsEqualsToPrimitive(t *testing.T) {
	That{false, t}.IsEqualsTo(false)
	That{true, t}.IsEqualsTo(true)

	That{0, t}.IsEqualsTo(0)
	That{1, t}.IsEqualsTo(1)
	That{-1, t}.IsEqualsTo(-1)

	That{int32(0), t}.IsEqualsTo(int32(0))
	That{int32(1), t}.IsEqualsTo(int32(1))
	That{int32(-1), t}.IsEqualsTo(int32(-1))

	That{int64(0), t}.IsEqualsTo(int64(0))
	That{int64(1), t}.IsEqualsTo(int64(1))
	That{int64(-1), t}.IsEqualsTo(int64(-1))

	That{uint(0), t}.IsEqualsTo(uint(0))
	That{uint(1), t}.IsEqualsTo(uint(1))

	That{uint32(0), t}.IsEqualsTo(uint32(0))
	That{uint32(1), t}.IsEqualsTo(uint32(1))

	That{uint64(0), t}.IsEqualsTo(uint64(0))
	That{uint64(1), t}.IsEqualsTo(uint64(1))

	That{float32(0.25), t}.IsEqualsTo(float32(0.25))
	That{float32(1.8), t}.IsEqualsTo(float32(1.8))

	That{float64(0.25), t}.IsEqualsTo(float64(0.25))
	That{float64(1.8), t}.IsEqualsTo(float64(1.8))

	That{"", t}.IsEqualsTo("")
	That{"abcXYZ", t}.IsEqualsTo("abcXYZ")
}
