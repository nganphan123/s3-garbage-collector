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

package filters_test

import (
	"cmp"
	"context"
	"slices"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/nganphan123/s3-garbage-collector/pkg/filters"
	"github.com/nganphan123/s3-garbage-collector/pkg/types"
	"github.com/stretchr/testify/assert"
)

func TestFilterByObjectKey(t *testing.T) {
	testCases := []struct {
		name string

		config    func() types.DeleteConfig
		s3Objects []types.S3Object

		expectedObjs  []types.S3Object
		expectedError string
	}{
		{
			name: "should return empty if selectors are empty",
			config: func() types.DeleteConfig {
				config := validDeleteConfig()
				config.Selectors = nil
				return config
			},
			s3Objects:    validS3Objects(),
			expectedObjs: nil,
		},
		{
			name:      "should return filtered objects as expected",
			config:    validDeleteConfig,
			s3Objects: validS3Objects(),
			expectedObjs: []types.S3Object{
				{
					Object: &s3types.Object{
						Key: aws.String("/storage/my-file"),
					},
					Tags: make(types.AWSTags),
				},
				{
					Object: &s3types.Object{
						Key: aws.String("/storage/file.pdf"),
					},
					Tags: make(types.AWSTags),
				},
				{
					Object: &s3types.Object{
						Key: aws.String("tmp-file.txt"),
					},
					Tags: make(types.AWSTags),
				},
			},
		},
		{
			name: "should return error if match expression is invalid/unsupported",
			config: func() types.DeleteConfig {
				config := validDeleteConfig()
				config.Selectors = append(config.Selectors, types.Selector{MatchExpression: "(?!)"})
				return config
			},
			s3Objects:     validS3Objects(),
			expectedError: `error parsing regexp: invalid or unsupported Perl syntax`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := types.FilterInput{
				Config: tc.config(),
				Objs:   tc.s3Objects,
			}

			output, err := objectKeyFilter().Filter(context.TODO(), input)
			if len(tc.expectedError) > 0 {
				assert.Error(t, err)
				assert.Regexp(t, tc.expectedError, err.Error())
			} else {
				assert.NoError(t, err)
			}

			// Sort object slices for comparing
			sortFn := func(a, b types.S3Object) int {
				return cmp.Compare(aws.ToString(a.Key), aws.ToString(b.Key))
			}
			slices.SortFunc(output.Objs, sortFn)
			slices.SortFunc(tc.expectedObjs, sortFn)

			assert.Truef(t, types.IsEqualS3ObjectSlice(output.Objs, tc.expectedObjs),
				"expected %v, but got %v", tc.expectedObjs, output.Objs)
		})
	}
}

// objectKeyFilter returns a pointer to an ObjectKeyFilter.
func objectKeyFilter() *filters.ObjectKeyFilter {
	return &filters.ObjectKeyFilter{}
}

// validDeleteConfig returns a DeleteConfig where each entry
// defines a matchExpression.
func validDeleteConfig() types.DeleteConfig {
	return types.DeleteConfig{
		ApiVersion: "v1alpha1",
		Kind:       "DeleteConfig",
		Selectors: []types.Selector{
			{MatchExpression: "\\.pdf$"},
			{MatchExpression: "^tmp"},
			{MatchExpression: "my-file"},
		},
	}
}

// validS3Objects returns a collection of "existing" S3 objects
// Note: These are partial structs where only Key and Tags are defined.
func validS3Objects() []types.S3Object {
	return []types.S3Object{
		{
			Object: &s3types.Object{
				Key: aws.String("/storage/my-file"),
			},
			Tags: make(types.AWSTags),
		},
		{
			Object: &s3types.Object{
				Key: aws.String("/storage/file.pdf"),
			},
			Tags: make(types.AWSTags),
		},
		{
			Object: &s3types.Object{
				Key: aws.String("/storage/tmp-file.txt"),
			},
			Tags: make(types.AWSTags),
		},
		{
			Object: &s3types.Object{
				Key: aws.String("/storage/important-secret.txt"),
			},
			Tags: make(types.AWSTags),
		},
		{
			Object: &s3types.Object{
				Key: aws.String("tmp-file.txt"),
			},
			Tags: make(types.AWSTags),
		},
		{
			Object: &s3types.Object{
				Key: aws.String("/etc/resolv.conf"),
			},
			Tags: make(types.AWSTags),
		},
	}
}
