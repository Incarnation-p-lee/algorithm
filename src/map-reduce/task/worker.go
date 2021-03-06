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
	"net/rpc"
)

// WorkerType indicates the type of worker, can be Map or Reduce worker.
type WorkerType int32

// The workerType Map and Reduce, will add more in future.
const (
	Map = iota
	Reduce
)

const (
	registryMethod = "MasterOperation.Register"
)

// Worker represent one worker of host machine from cluster.
type Worker struct {
	Type WorkerType
	ID   int32
}

// Register Worker to remote address, and update some data from Master node.
func (worker *Worker) Register(addr, path string) {
	client, err := rpc.DialHTTPPath("tcp", addr, path)
	if err != nil {
		log.Fatalf("Failed to dial address %s, %s", addr, err)
	}

	var response MasterResponse
	err = client.Call(registryMethod, *worker, &response)

	if err != nil {
		log.Fatalf("Failed to register to master: %s, %s", addr, err)
	}

	worker.ID = response.ID
}
