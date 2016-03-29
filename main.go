// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/yarpc/crossdock/execute"
	"github.com/yarpc/crossdock/output"
	"github.com/yarpc/crossdock/plan"
)

func main() {
	fmt.Printf("\nCrossdock starting...\n\n")
	plan := plan.New(plan.ReadConfigFromEnviron())

	fmt.Printf("Waiting on CROSSDOCK_CLIENTS=%v\n\n", plan.Config.Clients)
	execute.Wait(plan.Config.Clients, time.Duration(30)*time.Second)

	fmt.Printf("\nExecuting Matrix...\n\n")
	results := execute.Run(plan)

	if err := output.Stream(results); err != nil {
		fmt.Printf("\nTests did not pass!\n\n")
		os.Exit(1)
	}
	fmt.Printf("\nTests passed!\n\n")
}
