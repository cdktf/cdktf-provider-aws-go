// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrlifecyclepolicydocument


type DataAwsEcrLifecyclePolicyDocumentRuleSelection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#count_number DataAwsEcrLifecyclePolicyDocument#count_number}.
	CountNumber *float64 `field:"required" json:"countNumber" yaml:"countNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#count_type DataAwsEcrLifecyclePolicyDocument#count_type}.
	CountType *string `field:"required" json:"countType" yaml:"countType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#tag_status DataAwsEcrLifecyclePolicyDocument#tag_status}.
	TagStatus *string `field:"required" json:"tagStatus" yaml:"tagStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#count_unit DataAwsEcrLifecyclePolicyDocument#count_unit}.
	CountUnit *string `field:"optional" json:"countUnit" yaml:"countUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#tag_pattern_list DataAwsEcrLifecyclePolicyDocument#tag_pattern_list}.
	TagPatternList *[]*string `field:"optional" json:"tagPatternList" yaml:"tagPatternList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/data-sources/ecr_lifecycle_policy_document#tag_prefix_list DataAwsEcrLifecyclePolicyDocument#tag_prefix_list}.
	TagPrefixList *[]*string `field:"optional" json:"tagPrefixList" yaml:"tagPrefixList"`
}

