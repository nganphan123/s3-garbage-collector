// MIT License

// Copyright (c) 2025 Ngan Phan

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package types

import (
	"context"
)

// Filter is the base interface that each filter handler
// should implement.
type Filter interface {
	// Name provides the user-friendly name of the filter.
	Name() string

	// Filter returns the S3 objects that satisfy the matching criteria
	Filter(ctx context.Context, input FilterInput) (FilterOutput, error)
}

// FilterInput collects arguments to pass to the Filter handler.
type FilterInput struct {
	// Objs is a collection of S3 objects to be filtered.
	Objs []S3Object
	// Config is the deletion configurations, provided by the users.
	Config DeleteConfig
}

// FilterOutput collects the output of the Filter handler.
type FilterOutput struct {
	// Objs is a collection of filtered objects.
	Objs []S3Object
}
