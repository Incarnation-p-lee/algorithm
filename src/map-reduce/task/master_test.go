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

package task

import (
	"assert"
	"sort"
	"testing"
	"time"
	// "log"
)

const (
	delay time.Duration = 500 * time.Millisecond
)

func TestMasterListen(t *testing.T) {
	var data = []struct {
		expect int32
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
	}

	resetWorkerID()

	addr, path, master := "localhost:8000", "/TestMasterListen", new(Master)

	go master.Listen(addr, path)
	time.Sleep(delay) // Make sure rpc server is ready.

	assert.That{master.addr, t}.IsEqualsTo(addr)

	for _, d := range data {
		worker := new(Worker)
		worker.Register(addr, path)
		actual := worker.ID

		assert.That{actual, t}.IsEqualsTo(d.expect)
	}
}

func TestMasterListenParallel(t *testing.T) {
	var data = []struct {
		expect int
	}{
		{1},
		{2},
		{3},
		{4},
		{5},
		{6},
		{7},
		{8},
		{9},
	}

	resetWorkerID()

	addr, path, master := "localhost:8001", "/TestMasterListenParallel", new(Master)

	go master.Listen(addr, path)
	time.Sleep(delay) // Make sure rpc server is ready.

	assert.That{master.addr, t}.IsEqualsTo(addr)

	var workers []*Worker

	for i := 0; i < len(data); i++ {
		w := new(Worker)
		workers = append(workers, w)

		go w.Register(addr, path)
	}

	time.Sleep(delay) // Make sure all work registered

	var ids []int

	for _, w := range workers {
		ids = append(ids, int(w.ID))
	}

	sort.Ints(ids)

	for i, id := range ids {
		assert.That{id, t}.IsEqualsTo(data[i].expect)
	}
}
