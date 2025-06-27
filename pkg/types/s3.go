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
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3Object represents an object in the S3 bucket
// with additional details such as tags.
type S3Object struct {
	*s3types.Object `json:",inline"`
	Tags            AWSTags `json:"tags,omitempty"`
}

// AWSTagsFromSDKType converts a list of SDK-typed tags (slice)
// into AWSTags (map).
func S3ObjectFromSDKType(obj *s3types.Object, tags []s3types.Tag) *S3Object {
	return &S3Object{
		Object: obj,
		Tags:   AWSTagsFromSDKType(tags),
	}
}

// IsEqual returns whether the S3Object is equal to the
// one provided as input.
func (obj *S3Object) IsEqual(other *S3Object) bool {
	if obj == nil && other == nil {
		return true
	}

	if (obj == nil && other != nil) || (obj != nil && other == nil) {
		return false
	}

	return obj == other || aws.ToString(obj.Key) == aws.ToString(other.Key)
}

// String implements the Stringer interface for custom print output.
func (obj S3Object) String() string {
	return fmt.Sprintf("key: %s, tags: %v, lastModified: %v", aws.ToString(obj.Key), obj.Tags, obj.LastModified)
}

// IsEqualS3ObjectSlice returns whether slices of S3Objects are equal.
func IsEqualS3ObjectSlice(a, b []S3Object) bool {
	if len(a) != len(b) {
		return false
	}
	for i, obj := range a {
		expectedObj := &b[i]
		if !obj.IsEqual(expectedObj) {
			return false
		}
	}
	return true
}
