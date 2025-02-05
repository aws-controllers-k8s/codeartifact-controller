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

type AllowPublish string

const (
	AllowPublish_ALLOW AllowPublish = "ALLOW"
	AllowPublish_BLOCK AllowPublish = "BLOCK"
)

type AllowUpstream string

const (
	AllowUpstream_ALLOW AllowUpstream = "ALLOW"
	AllowUpstream_BLOCK AllowUpstream = "BLOCK"
)

type DomainStatus_SDK string

const (
	DomainStatus_SDK_Active  DomainStatus_SDK = "Active"
	DomainStatus_SDK_Deleted DomainStatus_SDK = "Deleted"
)

type EndpointType string

const (
	EndpointType_dualstack EndpointType = "dualstack"
	EndpointType_ipv4      EndpointType = "ipv4"
)

type ExternalConnectionStatus string

const (
	ExternalConnectionStatus_Available ExternalConnectionStatus = "Available"
)

type HashAlgorithm string

const (
	HashAlgorithm_MD5     HashAlgorithm = "MD5"
	HashAlgorithm_SHA_1   HashAlgorithm = "SHA-1"
	HashAlgorithm_SHA_256 HashAlgorithm = "SHA-256"
	HashAlgorithm_SHA_512 HashAlgorithm = "SHA-512"
)

type PackageFormat string

const (
	PackageFormat_cargo   PackageFormat = "cargo"
	PackageFormat_generic PackageFormat = "generic"
	PackageFormat_maven   PackageFormat = "maven"
	PackageFormat_npm     PackageFormat = "npm"
	PackageFormat_nuget   PackageFormat = "nuget"
	PackageFormat_pypi    PackageFormat = "pypi"
	PackageFormat_ruby    PackageFormat = "ruby"
	PackageFormat_swift   PackageFormat = "swift"
)

type PackageGroupAllowedRepositoryUpdateType string

const (
	PackageGroupAllowedRepositoryUpdateType_ADDED   PackageGroupAllowedRepositoryUpdateType = "ADDED"
	PackageGroupAllowedRepositoryUpdateType_REMOVED PackageGroupAllowedRepositoryUpdateType = "REMOVED"
)

type PackageGroupAssociationType string

const (
	PackageGroupAssociationType_STRONG PackageGroupAssociationType = "STRONG"
	PackageGroupAssociationType_WEAK   PackageGroupAssociationType = "WEAK"
)

type PackageGroupOriginRestrictionMode string

const (
	PackageGroupOriginRestrictionMode_ALLOW                       PackageGroupOriginRestrictionMode = "ALLOW"
	PackageGroupOriginRestrictionMode_ALLOW_SPECIFIC_REPOSITORIES PackageGroupOriginRestrictionMode = "ALLOW_SPECIFIC_REPOSITORIES"
	PackageGroupOriginRestrictionMode_BLOCK                       PackageGroupOriginRestrictionMode = "BLOCK"
	PackageGroupOriginRestrictionMode_INHERIT                     PackageGroupOriginRestrictionMode = "INHERIT"
)

type PackageGroupOriginRestrictionType string

const (
	PackageGroupOriginRestrictionType_EXTERNAL_UPSTREAM PackageGroupOriginRestrictionType = "EXTERNAL_UPSTREAM"
	PackageGroupOriginRestrictionType_INTERNAL_UPSTREAM PackageGroupOriginRestrictionType = "INTERNAL_UPSTREAM"
	PackageGroupOriginRestrictionType_PUBLISH           PackageGroupOriginRestrictionType = "PUBLISH"
)

type PackageVersionErrorCode string

const (
	PackageVersionErrorCode_ALREADY_EXISTS      PackageVersionErrorCode = "ALREADY_EXISTS"
	PackageVersionErrorCode_MISMATCHED_REVISION PackageVersionErrorCode = "MISMATCHED_REVISION"
	PackageVersionErrorCode_MISMATCHED_STATUS   PackageVersionErrorCode = "MISMATCHED_STATUS"
	PackageVersionErrorCode_NOT_ALLOWED         PackageVersionErrorCode = "NOT_ALLOWED"
	PackageVersionErrorCode_NOT_FOUND           PackageVersionErrorCode = "NOT_FOUND"
	PackageVersionErrorCode_SKIPPED             PackageVersionErrorCode = "SKIPPED"
)

type PackageVersionOriginType string

const (
	PackageVersionOriginType_EXTERNAL PackageVersionOriginType = "EXTERNAL"
	PackageVersionOriginType_INTERNAL PackageVersionOriginType = "INTERNAL"
	PackageVersionOriginType_UNKNOWN  PackageVersionOriginType = "UNKNOWN"
)

type PackageVersionSortType string

const (
	PackageVersionSortType_PUBLISHED_TIME PackageVersionSortType = "PUBLISHED_TIME"
)

type PackageVersionStatus string

const (
	PackageVersionStatus_Archived   PackageVersionStatus = "Archived"
	PackageVersionStatus_Deleted    PackageVersionStatus = "Deleted"
	PackageVersionStatus_Disposed   PackageVersionStatus = "Disposed"
	PackageVersionStatus_Published  PackageVersionStatus = "Published"
	PackageVersionStatus_Unfinished PackageVersionStatus = "Unfinished"
	PackageVersionStatus_Unlisted   PackageVersionStatus = "Unlisted"
)

type ResourceType string

const (
	ResourceType_asset           ResourceType = "asset"
	ResourceType_domain          ResourceType = "domain"
	ResourceType_package         ResourceType = "package"
	ResourceType_package_version ResourceType = "package-version"
	ResourceType_repository      ResourceType = "repository"
)

type ValidationExceptionReason string

const (
	ValidationExceptionReason_CANNOT_PARSE            ValidationExceptionReason = "CANNOT_PARSE"
	ValidationExceptionReason_ENCRYPTION_KEY_ERROR    ValidationExceptionReason = "ENCRYPTION_KEY_ERROR"
	ValidationExceptionReason_FIELD_VALIDATION_FAILED ValidationExceptionReason = "FIELD_VALIDATION_FAILED"
	ValidationExceptionReason_OTHER                   ValidationExceptionReason = "OTHER"
	ValidationExceptionReason_UNKNOWN_OPERATION       ValidationExceptionReason = "UNKNOWN_OPERATION"
)
