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

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PackageGroupSpec defines the desired state of PackageGroup.
type PackageGroupSpec struct {

	// The contact information for the created package group.
	//
	// Regex Pattern: `^\P{C}*$`
	ContactInfo *string `json:"contactInfo,omitempty"`
	// A description of the package group.
	//
	// Regex Pattern: `^\P{C}*$`
	Description *string `json:"description,omitempty"`
	// The name of the domain in which you want to create a package group.
	//
	// Regex Pattern: `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`
	// +kubebuilder:validation:Required
	Domain *string `json:"domain"`
	// The 12-digit account number of the Amazon Web Services account that owns
	// the domain. It does not include dashes or spaces.
	//
	// Regex Pattern: `^[0-9]{12}$`
	DomainOwner *string `json:"domainOwner,omitempty"`
	// The pattern of the package group to create. The pattern is also the identifier
	// of the package group.
	//
	// Regex Pattern: `^[^\p{C}\p{IsWhitespace}]+$`
	// +kubebuilder:validation:Required
	Pattern *string `json:"pattern"`
	// One or more tag key-value pairs for the package group.
	Tags []*Tag `json:"tags,omitempty"`
}

// PackageGroupStatus defines the observed state of PackageGroup
type PackageGroupStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRs managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// A timestamp that represents the date and time the package group was created.
	// +kubebuilder:validation:Optional
	CreatedTime *metav1.Time `json:"createdTime,omitempty"`
	// The name of the domain that contains the package group.
	//
	// Regex Pattern: `^[a-z][a-z0-9\-]{0,48}[a-z0-9]$`
	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty"`
	// The package group origin configuration that determines how package versions
	// can enter repositories.
	// +kubebuilder:validation:Optional
	OriginConfiguration *PackageGroupOriginConfiguration `json:"originConfiguration,omitempty"`
	// The direct parent package group of the package group.
	// +kubebuilder:validation:Optional
	Parent *PackageGroupReference `json:"parent,omitempty"`
}

// PackageGroup is the Schema for the PackageGroups API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type PackageGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PackageGroupSpec   `json:"spec,omitempty"`
	Status            PackageGroupStatus `json:"status,omitempty"`
}

// PackageGroupList contains a list of PackageGroup
// +kubebuilder:object:root=true
type PackageGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PackageGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PackageGroup{}, &PackageGroupList{})
}
