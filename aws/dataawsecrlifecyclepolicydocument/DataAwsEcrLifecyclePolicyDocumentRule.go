// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrlifecyclepolicydocument


type DataAwsEcrLifecyclePolicyDocumentRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/ecr_lifecycle_policy_document#priority DataAwsEcrLifecyclePolicyDocument#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/ecr_lifecycle_policy_document#action DataAwsEcrLifecyclePolicyDocument#action}
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/ecr_lifecycle_policy_document#description DataAwsEcrLifecyclePolicyDocument#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// selection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/ecr_lifecycle_policy_document#selection DataAwsEcrLifecyclePolicyDocument#selection}
	Selection interface{} `field:"optional" json:"selection" yaml:"selection"`
}

