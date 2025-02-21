// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lambdacodesigningconfig


type LambdaCodeSigningConfigAllowedPublishers struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/lambda_code_signing_config#signing_profile_version_arns LambdaCodeSigningConfig#signing_profile_version_arns}.
	SigningProfileVersionArns *[]*string `field:"required" json:"signingProfileVersionArns" yaml:"signingProfileVersionArns"`
}

