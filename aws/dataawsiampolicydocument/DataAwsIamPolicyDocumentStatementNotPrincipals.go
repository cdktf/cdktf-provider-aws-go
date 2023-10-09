// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsiampolicydocument


type DataAwsIamPolicyDocumentStatementNotPrincipals struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/data-sources/iam_policy_document#identifiers DataAwsIamPolicyDocument#identifiers}.
	Identifiers *[]*string `field:"required" json:"identifiers" yaml:"identifiers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/data-sources/iam_policy_document#type DataAwsIamPolicyDocument#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

