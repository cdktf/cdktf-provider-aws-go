// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiampolicydocument


type DataAwsIamPolicyDocumentStatementCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/iam_policy_document#test DataAwsIamPolicyDocument#test}.
	Test *string `field:"required" json:"test" yaml:"test"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/iam_policy_document#values DataAwsIamPolicyDocument#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/data-sources/iam_policy_document#variable DataAwsIamPolicyDocument#variable}.
	Variable *string `field:"required" json:"variable" yaml:"variable"`
}

