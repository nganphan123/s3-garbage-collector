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

package filters

import (
	"context"
	"regexp"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/nganphan123/s3-garbage-collector/pkg/types"
)

type ObjectKeyFilter struct{}

var _ types.Filter = (*ObjectKeyFilter)(nil)

func (f *ObjectKeyFilter) Name() string {
	return "Filter by object key"
}

func (f *ObjectKeyFilter) Filter(ctx context.Context, input types.FilterInput) (types.FilterOutput, error) {
	output := types.FilterOutput{}
	selectors := input.Config.Selectors

	// If no selectors are defined, no objects (S3 files)
	// will be selected (i.e. return empty).
	if len(selectors) == 0 {
		return output, nil
	}

	matchExprs := make([]string, 0)
	for _, selector := range selectors {
		if len(selector.MatchExpression) > 0 {
			matchExprs = append(matchExprs, selector.MatchExpression)
		}
	}

	regs, err := f.BuildRegex(matchExprs)
	if err != nil {
		return output, err
	}

	for _, obj := range input.Objs {
		for _, reg := range regs {
			if reg.MatchString(aws.ToString(obj.Key)) {
				output.Objs = append(output.Objs, obj)
				break
			}
		}
	}

	return output, nil
}

// BuildRegex returns a collection of compiled regular expressions
// from a collection of input expressions.
func (f *ObjectKeyFilter) BuildRegex(matchExprs []string) ([]*regexp.Regexp, error) {
	regs := make([]*regexp.Regexp, 0)
	for _, matchExpr := range matchExprs {
		regex, err := regexp.Compile(matchExpr)
		if err != nil {
			return regs, err
		}
		regs = append(regs, regex)
	}
	return regs, nil
}
