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
	"log"
	"net/http"
	"net/rpc"
	// assert "github.com/attic-labs/noms/go/d"
)

// Master represent the Master node as rpc server.
type Master struct {
	addr      string
	operation *MasterOperation
}

// Listen will register MasterOperation as rpc provide operation from
// Master node, and listen the given address.
func (master *Master) Listen(addr string) {
	op := new(MasterOperation)

	rpc.Register(op)
	rpc.HandleHTTP()

	master.addr = addr
	master.operation = op

	log.Fatal(http.ListenAndServe(addr, nil))
}
