// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetOperationPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#failure_tolerance_count CloudformationStackSet#failure_tolerance_count}.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#failure_tolerance_percentage CloudformationStackSet#failure_tolerance_percentage}.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#max_concurrent_count CloudformationStackSet#max_concurrent_count}.
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#max_concurrent_percentage CloudformationStackSet#max_concurrent_percentage}.
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#region_concurrency_type CloudformationStackSet#region_concurrency_type}.
	RegionConcurrencyType *string `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/cloudformation_stack_set#region_order CloudformationStackSet#region_order}.
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

