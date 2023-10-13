// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawskmssecret


type DataAwsKmsSecretSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/kms_secret#name DataAwsKmsSecret#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/kms_secret#payload DataAwsKmsSecret#payload}.
	Payload *string `field:"required" json:"payload" yaml:"payload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/kms_secret#context DataAwsKmsSecret#context}.
	Context *map[string]*string `field:"optional" json:"context" yaml:"context"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/data-sources/kms_secret#grant_tokens DataAwsKmsSecret#grant_tokens}.
	GrantTokens *[]*string `field:"optional" json:"grantTokens" yaml:"grantTokens"`
}

