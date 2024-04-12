// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiampolicydocument


type DataAwsIamPolicyDocumentStatement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#actions DataAwsIamPolicyDocument#actions}.
	Actions *[]*string `field:"optional" json:"actions" yaml:"actions"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#condition DataAwsIamPolicyDocument#condition}
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#effect DataAwsIamPolicyDocument#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#not_actions DataAwsIamPolicyDocument#not_actions}.
	NotActions *[]*string `field:"optional" json:"notActions" yaml:"notActions"`
	// not_principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#not_principals DataAwsIamPolicyDocument#not_principals}
	NotPrincipals interface{} `field:"optional" json:"notPrincipals" yaml:"notPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#not_resources DataAwsIamPolicyDocument#not_resources}.
	NotResources *[]*string `field:"optional" json:"notResources" yaml:"notResources"`
	// principals block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#principals DataAwsIamPolicyDocument#principals}
	Principals interface{} `field:"optional" json:"principals" yaml:"principals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#resources DataAwsIamPolicyDocument#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/data-sources/iam_policy_document#sid DataAwsIamPolicyDocument#sid}.
	Sid *string `field:"optional" json:"sid" yaml:"sid"`
}

