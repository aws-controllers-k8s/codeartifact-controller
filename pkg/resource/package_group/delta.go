// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package package_group

import (
	"bytes"
	"reflect"

	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	acktags "github.com/aws-controllers-k8s/runtime/pkg/tags"
)

// Hack to avoid import errors during build...
var (
	_ = &bytes.Buffer{}
	_ = &reflect.Method{}
	_ = &acktags.Tags{}
)

// newResourceDelta returns a new `ackcompare.Delta` used to compare two
// resources
func newResourceDelta(
	a *resource,
	b *resource,
) *ackcompare.Delta {
	delta := ackcompare.NewDelta()
	if (a == nil && b != nil) ||
		(a != nil && b == nil) {
		delta.Add("", a, b)
		return delta
	}
	compareTags(delta, a, b)

	if ackcompare.HasNilDifference(a.ko.Spec.ContactInfo, b.ko.Spec.ContactInfo) {
		delta.Add("Spec.ContactInfo", a.ko.Spec.ContactInfo, b.ko.Spec.ContactInfo)
	} else if a.ko.Spec.ContactInfo != nil && b.ko.Spec.ContactInfo != nil {
		if *a.ko.Spec.ContactInfo != *b.ko.Spec.ContactInfo {
			delta.Add("Spec.ContactInfo", a.ko.Spec.ContactInfo, b.ko.Spec.ContactInfo)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Description, b.ko.Spec.Description) {
		delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
	} else if a.ko.Spec.Description != nil && b.ko.Spec.Description != nil {
		if *a.ko.Spec.Description != *b.ko.Spec.Description {
			delta.Add("Spec.Description", a.ko.Spec.Description, b.ko.Spec.Description)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Domain, b.ko.Spec.Domain) {
		delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
	} else if a.ko.Spec.Domain != nil && b.ko.Spec.Domain != nil {
		if *a.ko.Spec.Domain != *b.ko.Spec.Domain {
			delta.Add("Spec.Domain", a.ko.Spec.Domain, b.ko.Spec.Domain)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.DomainOwner, b.ko.Spec.DomainOwner) {
		delta.Add("Spec.DomainOwner", a.ko.Spec.DomainOwner, b.ko.Spec.DomainOwner)
	} else if a.ko.Spec.DomainOwner != nil && b.ko.Spec.DomainOwner != nil {
		if *a.ko.Spec.DomainOwner != *b.ko.Spec.DomainOwner {
			delta.Add("Spec.DomainOwner", a.ko.Spec.DomainOwner, b.ko.Spec.DomainOwner)
		}
	}
	if ackcompare.HasNilDifference(a.ko.Spec.Pattern, b.ko.Spec.Pattern) {
		delta.Add("Spec.Pattern", a.ko.Spec.Pattern, b.ko.Spec.Pattern)
	} else if a.ko.Spec.Pattern != nil && b.ko.Spec.Pattern != nil {
		if *a.ko.Spec.Pattern != *b.ko.Spec.Pattern {
			delta.Add("Spec.Pattern", a.ko.Spec.Pattern, b.ko.Spec.Pattern)
		}
	}

	return delta
}
