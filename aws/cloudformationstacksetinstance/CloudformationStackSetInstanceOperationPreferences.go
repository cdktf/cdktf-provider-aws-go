// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstacksetinstance


type CloudformationStackSetInstanceOperationPreferences struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#concurrency_mode CloudformationStackSetInstance#concurrency_mode}.
	ConcurrencyMode *string `field:"optional" json:"concurrencyMode" yaml:"concurrencyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#failure_tolerance_count CloudformationStackSetInstance#failure_tolerance_count}.
	FailureToleranceCount *float64 `field:"optional" json:"failureToleranceCount" yaml:"failureToleranceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#failure_tolerance_percentage CloudformationStackSetInstance#failure_tolerance_percentage}.
	FailureTolerancePercentage *float64 `field:"optional" json:"failureTolerancePercentage" yaml:"failureTolerancePercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#max_concurrent_count CloudformationStackSetInstance#max_concurrent_count}.
	MaxConcurrentCount *float64 `field:"optional" json:"maxConcurrentCount" yaml:"maxConcurrentCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#max_concurrent_percentage CloudformationStackSetInstance#max_concurrent_percentage}.
	MaxConcurrentPercentage *float64 `field:"optional" json:"maxConcurrentPercentage" yaml:"maxConcurrentPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#region_concurrency_type CloudformationStackSetInstance#region_concurrency_type}.
	RegionConcurrencyType *string `field:"optional" json:"regionConcurrencyType" yaml:"regionConcurrencyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/cloudformation_stack_set_instance#region_order CloudformationStackSetInstance#region_order}.
	RegionOrder *[]*string `field:"optional" json:"regionOrder" yaml:"regionOrder"`
}

