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
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	ackcondition "github.com/aws-controllers-k8s/runtime/pkg/condition"
	ackerr "github.com/aws-controllers-k8s/runtime/pkg/errors"
	ackrequeue "github.com/aws-controllers-k8s/runtime/pkg/requeue"
	ackrtlog "github.com/aws-controllers-k8s/runtime/pkg/runtime/log"
	"github.com/aws/aws-sdk-go-v2/aws"
	svcsdk "github.com/aws/aws-sdk-go-v2/service/codeartifact"
	svcsdktypes "github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/aws/smithy-go"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	svcapitypes "github.com/aws-controllers-k8s/codeartifact-controller/apis/v1alpha1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
	_ = strings.ToLower("")
	_ = &svcsdk.Client{}
	_ = &svcapitypes.PackageGroup{}
	_ = ackv1alpha1.AWSAccountID("")
	_ = &ackerr.NotFound
	_ = &ackcondition.NotManagedMessage
	_ = &reflect.Value{}
	_ = fmt.Sprintf("")
	_ = &ackrequeue.NoRequeue{}
	_ = &aws.Config{}
)

// sdkFind returns SDK-specific information about a supplied resource
func (rm *resourceManager) sdkFind(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkFind")
	defer func() {
		exit(err)
	}()
	// If any required fields in the input shape are missing, AWS resource is
	// not created yet. Return NotFound here to indicate to callers that the
	// resource isn't yet created.
	if rm.requiredFieldsMissingFromReadOneInput(r) {
		return nil, ackerr.NotFound
	}

	input, err := rm.newDescribeRequestPayload(r)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.DescribePackageGroupOutput
	resp, err = rm.sdkapi.DescribePackageGroup(ctx, input)
	rm.metrics.RecordAPICall("READ_ONE", "DescribePackageGroup", err)
	if err != nil {
		var awsErr smithy.APIError
		if errors.As(err, &awsErr) && awsErr.ErrorCode() == "ResourceNotFoundException" {
			return nil, ackerr.NotFound
		}
		return nil, err
	}

	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := r.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.PackageGroup.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.PackageGroup.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.PackageGroup.ContactInfo != nil {
		ko.Spec.ContactInfo = resp.PackageGroup.ContactInfo
	} else {
		ko.Spec.ContactInfo = nil
	}
	if resp.PackageGroup.CreatedTime != nil {
		ko.Status.CreatedTime = &metav1.Time{*resp.PackageGroup.CreatedTime}
	} else {
		ko.Status.CreatedTime = nil
	}
	if resp.PackageGroup.Description != nil {
		ko.Spec.Description = resp.PackageGroup.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.PackageGroup.DomainName != nil {
		ko.Status.DomainName = resp.PackageGroup.DomainName
	} else {
		ko.Status.DomainName = nil
	}
	if resp.PackageGroup.DomainOwner != nil {
		ko.Spec.DomainOwner = resp.PackageGroup.DomainOwner
	} else {
		ko.Spec.DomainOwner = nil
	}
	if resp.PackageGroup.OriginConfiguration != nil {
		f6 := &svcapitypes.PackageGroupOriginConfiguration{}
		if resp.PackageGroup.OriginConfiguration.Restrictions != nil {
			f6f0 := map[string]*svcapitypes.PackageGroupOriginRestriction{}
			for f6f0key, f6f0valiter := range resp.PackageGroup.OriginConfiguration.Restrictions {
				f6f0val := &svcapitypes.PackageGroupOriginRestriction{}
				if f6f0valiter.EffectiveMode != "" {
					f6f0val.EffectiveMode = aws.String(string(f6f0valiter.EffectiveMode))
				}
				if f6f0valiter.InheritedFrom != nil {
					f6f0valf1 := &svcapitypes.PackageGroupReference{}
					if f6f0valiter.InheritedFrom.Arn != nil {
						f6f0valf1.ARN = f6f0valiter.InheritedFrom.Arn
					}
					if f6f0valiter.InheritedFrom.Pattern != nil {
						f6f0valf1.Pattern = f6f0valiter.InheritedFrom.Pattern
					}
					f6f0val.InheritedFrom = f6f0valf1
				}
				if f6f0valiter.Mode != "" {
					f6f0val.Mode = aws.String(string(f6f0valiter.Mode))
				}
				if f6f0valiter.RepositoriesCount != nil {
					f6f0val.RepositoriesCount = f6f0valiter.RepositoriesCount
				}
				f6f0[f6f0key] = f6f0val
			}
			f6.Restrictions = f6f0
		}
		ko.Status.OriginConfiguration = f6
	} else {
		ko.Status.OriginConfiguration = nil
	}
	if resp.PackageGroup.Parent != nil {
		f7 := &svcapitypes.PackageGroupReference{}
		if resp.PackageGroup.Parent.Arn != nil {
			f7.ARN = resp.PackageGroup.Parent.Arn
		}
		if resp.PackageGroup.Parent.Pattern != nil {
			f7.Pattern = resp.PackageGroup.Parent.Pattern
		}
		ko.Status.Parent = f7
	} else {
		ko.Status.Parent = nil
	}
	if resp.PackageGroup.Pattern != nil {
		ko.Spec.Pattern = resp.PackageGroup.Pattern
	} else {
		ko.Spec.Pattern = nil
	}

	rm.setStatusDefaults(ko)
	if err := rm.setResourceAdditionalFields(ctx, ko); err != nil {
		return nil, err
	}
	return &resource{ko}, nil
}

// requiredFieldsMissingFromReadOneInput returns true if there are any fields
// for the ReadOne Input shape that are required but not present in the
// resource's Spec or Status
func (rm *resourceManager) requiredFieldsMissingFromReadOneInput(
	r *resource,
) bool {
	return r.ko.Spec.Domain == nil || r.ko.Spec.Pattern == nil

}

// newDescribeRequestPayload returns SDK-specific struct for the HTTP request
// payload of the Describe API call for the resource
func (rm *resourceManager) newDescribeRequestPayload(
	r *resource,
) (*svcsdk.DescribePackageGroupInput, error) {
	res := &svcsdk.DescribePackageGroupInput{}

	if r.ko.Spec.Domain != nil {
		res.Domain = r.ko.Spec.Domain
	}
	if r.ko.Spec.DomainOwner != nil {
		res.DomainOwner = r.ko.Spec.DomainOwner
	}
	if r.ko.Spec.Pattern != nil {
		res.PackageGroup = r.ko.Spec.Pattern
	}

	return res, nil
}

// sdkCreate creates the supplied resource in the backend AWS service API and
// returns a copy of the resource with resource fields (in both Spec and
// Status) filled in with values from the CREATE API operation's Output shape.
func (rm *resourceManager) sdkCreate(
	ctx context.Context,
	desired *resource,
) (created *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkCreate")
	defer func() {
		exit(err)
	}()
	input, err := rm.newCreateRequestPayload(ctx, desired)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.CreatePackageGroupOutput
	_ = resp
	resp, err = rm.sdkapi.CreatePackageGroup(ctx, input)
	rm.metrics.RecordAPICall("CREATE", "CreatePackageGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.PackageGroup.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.PackageGroup.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.PackageGroup.ContactInfo != nil {
		ko.Spec.ContactInfo = resp.PackageGroup.ContactInfo
	} else {
		ko.Spec.ContactInfo = nil
	}
	if resp.PackageGroup.CreatedTime != nil {
		ko.Status.CreatedTime = &metav1.Time{*resp.PackageGroup.CreatedTime}
	} else {
		ko.Status.CreatedTime = nil
	}
	if resp.PackageGroup.Description != nil {
		ko.Spec.Description = resp.PackageGroup.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.PackageGroup.DomainName != nil {
		ko.Status.DomainName = resp.PackageGroup.DomainName
	} else {
		ko.Status.DomainName = nil
	}
	if resp.PackageGroup.DomainOwner != nil {
		ko.Spec.DomainOwner = resp.PackageGroup.DomainOwner
	} else {
		ko.Spec.DomainOwner = nil
	}
	if resp.PackageGroup.OriginConfiguration != nil {
		f6 := &svcapitypes.PackageGroupOriginConfiguration{}
		if resp.PackageGroup.OriginConfiguration.Restrictions != nil {
			f6f0 := map[string]*svcapitypes.PackageGroupOriginRestriction{}
			for f6f0key, f6f0valiter := range resp.PackageGroup.OriginConfiguration.Restrictions {
				f6f0val := &svcapitypes.PackageGroupOriginRestriction{}
				if f6f0valiter.EffectiveMode != "" {
					f6f0val.EffectiveMode = aws.String(string(f6f0valiter.EffectiveMode))
				}
				if f6f0valiter.InheritedFrom != nil {
					f6f0valf1 := &svcapitypes.PackageGroupReference{}
					if f6f0valiter.InheritedFrom.Arn != nil {
						f6f0valf1.ARN = f6f0valiter.InheritedFrom.Arn
					}
					if f6f0valiter.InheritedFrom.Pattern != nil {
						f6f0valf1.Pattern = f6f0valiter.InheritedFrom.Pattern
					}
					f6f0val.InheritedFrom = f6f0valf1
				}
				if f6f0valiter.Mode != "" {
					f6f0val.Mode = aws.String(string(f6f0valiter.Mode))
				}
				if f6f0valiter.RepositoriesCount != nil {
					f6f0val.RepositoriesCount = f6f0valiter.RepositoriesCount
				}
				f6f0[f6f0key] = f6f0val
			}
			f6.Restrictions = f6f0
		}
		ko.Status.OriginConfiguration = f6
	} else {
		ko.Status.OriginConfiguration = nil
	}
	if resp.PackageGroup.Parent != nil {
		f7 := &svcapitypes.PackageGroupReference{}
		if resp.PackageGroup.Parent.Arn != nil {
			f7.ARN = resp.PackageGroup.Parent.Arn
		}
		if resp.PackageGroup.Parent.Pattern != nil {
			f7.Pattern = resp.PackageGroup.Parent.Pattern
		}
		ko.Status.Parent = f7
	} else {
		ko.Status.Parent = nil
	}
	if resp.PackageGroup.Pattern != nil {
		ko.Spec.Pattern = resp.PackageGroup.Pattern
	} else {
		ko.Spec.Pattern = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newCreateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Create API call for the resource
func (rm *resourceManager) newCreateRequestPayload(
	ctx context.Context,
	r *resource,
) (*svcsdk.CreatePackageGroupInput, error) {
	res := &svcsdk.CreatePackageGroupInput{}

	if r.ko.Spec.ContactInfo != nil {
		res.ContactInfo = r.ko.Spec.ContactInfo
	}
	if r.ko.Spec.Description != nil {
		res.Description = r.ko.Spec.Description
	}
	if r.ko.Spec.Domain != nil {
		res.Domain = r.ko.Spec.Domain
	}
	if r.ko.Spec.DomainOwner != nil {
		res.DomainOwner = r.ko.Spec.DomainOwner
	}
	if r.ko.Spec.Pattern != nil {
		res.PackageGroup = r.ko.Spec.Pattern
	}
	if r.ko.Spec.Tags != nil {
		f5 := []svcsdktypes.Tag{}
		for _, f5iter := range r.ko.Spec.Tags {
			f5elem := &svcsdktypes.Tag{}
			if f5iter.Key != nil {
				f5elem.Key = f5iter.Key
			}
			if f5iter.Value != nil {
				f5elem.Value = f5iter.Value
			}
			f5 = append(f5, *f5elem)
		}
		res.Tags = f5
	}

	return res, nil
}

// sdkUpdate patches the supplied resource in the backend AWS service API and
// returns a new resource with updated fields.
func (rm *resourceManager) sdkUpdate(
	ctx context.Context,
	desired *resource,
	latest *resource,
	delta *ackcompare.Delta,
) (updated *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkUpdate")
	defer func() {
		exit(err)
	}()
	if delta.DifferentAt("Spec.Tags") {
		err = rm.syncTags(ctx, desired, latest)
		if err != nil {
			return nil, err
		}
	}

	if !delta.DifferentExcept("Spec.Tags") {
		return desired, nil
	}
	input, err := rm.newUpdateRequestPayload(ctx, desired, delta)
	if err != nil {
		return nil, err
	}

	var resp *svcsdk.UpdatePackageGroupOutput
	_ = resp
	resp, err = rm.sdkapi.UpdatePackageGroup(ctx, input)
	rm.metrics.RecordAPICall("UPDATE", "UpdatePackageGroup", err)
	if err != nil {
		return nil, err
	}
	// Merge in the information we read from the API call above to the copy of
	// the original Kubernetes object we passed to the function
	ko := desired.ko.DeepCopy()

	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if resp.PackageGroup.Arn != nil {
		arn := ackv1alpha1.AWSResourceName(*resp.PackageGroup.Arn)
		ko.Status.ACKResourceMetadata.ARN = &arn
	}
	if resp.PackageGroup.ContactInfo != nil {
		ko.Spec.ContactInfo = resp.PackageGroup.ContactInfo
	} else {
		ko.Spec.ContactInfo = nil
	}
	if resp.PackageGroup.CreatedTime != nil {
		ko.Status.CreatedTime = &metav1.Time{*resp.PackageGroup.CreatedTime}
	} else {
		ko.Status.CreatedTime = nil
	}
	if resp.PackageGroup.Description != nil {
		ko.Spec.Description = resp.PackageGroup.Description
	} else {
		ko.Spec.Description = nil
	}
	if resp.PackageGroup.DomainName != nil {
		ko.Status.DomainName = resp.PackageGroup.DomainName
	} else {
		ko.Status.DomainName = nil
	}
	if resp.PackageGroup.DomainOwner != nil {
		ko.Spec.DomainOwner = resp.PackageGroup.DomainOwner
	} else {
		ko.Spec.DomainOwner = nil
	}
	if resp.PackageGroup.OriginConfiguration != nil {
		f6 := &svcapitypes.PackageGroupOriginConfiguration{}
		if resp.PackageGroup.OriginConfiguration.Restrictions != nil {
			f6f0 := map[string]*svcapitypes.PackageGroupOriginRestriction{}
			for f6f0key, f6f0valiter := range resp.PackageGroup.OriginConfiguration.Restrictions {
				f6f0val := &svcapitypes.PackageGroupOriginRestriction{}
				if f6f0valiter.EffectiveMode != "" {
					f6f0val.EffectiveMode = aws.String(string(f6f0valiter.EffectiveMode))
				}
				if f6f0valiter.InheritedFrom != nil {
					f6f0valf1 := &svcapitypes.PackageGroupReference{}
					if f6f0valiter.InheritedFrom.Arn != nil {
						f6f0valf1.ARN = f6f0valiter.InheritedFrom.Arn
					}
					if f6f0valiter.InheritedFrom.Pattern != nil {
						f6f0valf1.Pattern = f6f0valiter.InheritedFrom.Pattern
					}
					f6f0val.InheritedFrom = f6f0valf1
				}
				if f6f0valiter.Mode != "" {
					f6f0val.Mode = aws.String(string(f6f0valiter.Mode))
				}
				if f6f0valiter.RepositoriesCount != nil {
					f6f0val.RepositoriesCount = f6f0valiter.RepositoriesCount
				}
				f6f0[f6f0key] = f6f0val
			}
			f6.Restrictions = f6f0
		}
		ko.Status.OriginConfiguration = f6
	} else {
		ko.Status.OriginConfiguration = nil
	}
	if resp.PackageGroup.Parent != nil {
		f7 := &svcapitypes.PackageGroupReference{}
		if resp.PackageGroup.Parent.Arn != nil {
			f7.ARN = resp.PackageGroup.Parent.Arn
		}
		if resp.PackageGroup.Parent.Pattern != nil {
			f7.Pattern = resp.PackageGroup.Parent.Pattern
		}
		ko.Status.Parent = f7
	} else {
		ko.Status.Parent = nil
	}
	if resp.PackageGroup.Pattern != nil {
		ko.Spec.Pattern = resp.PackageGroup.Pattern
	} else {
		ko.Spec.Pattern = nil
	}

	rm.setStatusDefaults(ko)
	return &resource{ko}, nil
}

// newUpdateRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Update API call for the resource
func (rm *resourceManager) newUpdateRequestPayload(
	ctx context.Context,
	r *resource,
	delta *ackcompare.Delta,
) (*svcsdk.UpdatePackageGroupInput, error) {
	res := &svcsdk.UpdatePackageGroupInput{}

	if r.ko.Spec.ContactInfo != nil {
		res.ContactInfo = r.ko.Spec.ContactInfo
	}
	if r.ko.Spec.Description != nil {
		res.Description = r.ko.Spec.Description
	}
	if r.ko.Spec.Domain != nil {
		res.Domain = r.ko.Spec.Domain
	}
	if r.ko.Spec.DomainOwner != nil {
		res.DomainOwner = r.ko.Spec.DomainOwner
	}
	if r.ko.Spec.Pattern != nil {
		res.PackageGroup = r.ko.Spec.Pattern
	}

	return res, nil
}

// sdkDelete deletes the supplied resource in the backend AWS service API
func (rm *resourceManager) sdkDelete(
	ctx context.Context,
	r *resource,
) (latest *resource, err error) {
	rlog := ackrtlog.FromContext(ctx)
	exit := rlog.Trace("rm.sdkDelete")
	defer func() {
		exit(err)
	}()
	input, err := rm.newDeleteRequestPayload(r)
	if err != nil {
		return nil, err
	}
	var resp *svcsdk.DeletePackageGroupOutput
	_ = resp
	resp, err = rm.sdkapi.DeletePackageGroup(ctx, input)
	rm.metrics.RecordAPICall("DELETE", "DeletePackageGroup", err)
	return nil, err
}

// newDeleteRequestPayload returns an SDK-specific struct for the HTTP request
// payload of the Delete API call for the resource
func (rm *resourceManager) newDeleteRequestPayload(
	r *resource,
) (*svcsdk.DeletePackageGroupInput, error) {
	res := &svcsdk.DeletePackageGroupInput{}

	if r.ko.Spec.Domain != nil {
		res.Domain = r.ko.Spec.Domain
	}
	if r.ko.Spec.DomainOwner != nil {
		res.DomainOwner = r.ko.Spec.DomainOwner
	}
	if r.ko.Spec.Pattern != nil {
		res.PackageGroup = r.ko.Spec.Pattern
	}

	return res, nil
}

// setStatusDefaults sets default properties into supplied custom resource
func (rm *resourceManager) setStatusDefaults(
	ko *svcapitypes.PackageGroup,
) {
	if ko.Status.ACKResourceMetadata == nil {
		ko.Status.ACKResourceMetadata = &ackv1alpha1.ResourceMetadata{}
	}
	if ko.Status.ACKResourceMetadata.Region == nil {
		ko.Status.ACKResourceMetadata.Region = &rm.awsRegion
	}
	if ko.Status.ACKResourceMetadata.OwnerAccountID == nil {
		ko.Status.ACKResourceMetadata.OwnerAccountID = &rm.awsAccountID
	}
	if ko.Status.Conditions == nil {
		ko.Status.Conditions = []*ackv1alpha1.Condition{}
	}
}

// updateConditions returns updated resource, true; if conditions were updated
// else it returns nil, false
func (rm *resourceManager) updateConditions(
	r *resource,
	onSuccess bool,
	err error,
) (*resource, bool) {
	ko := r.ko.DeepCopy()
	rm.setStatusDefaults(ko)

	// Terminal condition
	var terminalCondition *ackv1alpha1.Condition = nil
	var recoverableCondition *ackv1alpha1.Condition = nil
	var syncCondition *ackv1alpha1.Condition = nil
	for _, condition := range ko.Status.Conditions {
		if condition.Type == ackv1alpha1.ConditionTypeTerminal {
			terminalCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeRecoverable {
			recoverableCondition = condition
		}
		if condition.Type == ackv1alpha1.ConditionTypeResourceSynced {
			syncCondition = condition
		}
	}
	var termError *ackerr.TerminalError
	if rm.terminalAWSError(err) || err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
		if terminalCondition == nil {
			terminalCondition = &ackv1alpha1.Condition{
				Type: ackv1alpha1.ConditionTypeTerminal,
			}
			ko.Status.Conditions = append(ko.Status.Conditions, terminalCondition)
		}
		var errorMessage = ""
		if err == ackerr.SecretTypeNotSupported || err == ackerr.SecretNotFound || errors.As(err, &termError) {
			errorMessage = err.Error()
		} else {
			awsErr, _ := ackerr.AWSError(err)
			errorMessage = awsErr.Error()
		}
		terminalCondition.Status = corev1.ConditionTrue
		terminalCondition.Message = &errorMessage
	} else {
		// Clear the terminal condition if no longer present
		if terminalCondition != nil {
			terminalCondition.Status = corev1.ConditionFalse
			terminalCondition.Message = nil
		}
		// Handling Recoverable Conditions
		if err != nil {
			if recoverableCondition == nil {
				// Add a new Condition containing a non-terminal error
				recoverableCondition = &ackv1alpha1.Condition{
					Type: ackv1alpha1.ConditionTypeRecoverable,
				}
				ko.Status.Conditions = append(ko.Status.Conditions, recoverableCondition)
			}
			recoverableCondition.Status = corev1.ConditionTrue
			awsErr, _ := ackerr.AWSError(err)
			errorMessage := err.Error()
			if awsErr != nil {
				errorMessage = awsErr.Error()
			}
			recoverableCondition.Message = &errorMessage
		} else if recoverableCondition != nil {
			recoverableCondition.Status = corev1.ConditionFalse
			recoverableCondition.Message = nil
		}
	}
	// Required to avoid the "declared but not used" error in the default case
	_ = syncCondition
	if terminalCondition != nil || recoverableCondition != nil || syncCondition != nil {
		return &resource{ko}, true // updated
	}
	return nil, false // not updated
}

// terminalAWSError returns awserr, true; if the supplied error is an aws Error type
// and if the exception indicates that it is a Terminal exception
// 'Terminal' exception are specified in generator configuration
func (rm *resourceManager) terminalAWSError(err error) bool {
	// No terminal_errors specified for this resource in generator config
	return false
}
