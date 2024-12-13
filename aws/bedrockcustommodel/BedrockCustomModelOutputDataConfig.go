// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockcustommodel


type BedrockCustomModelOutputDataConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/bedrock_custom_model#s3_uri BedrockCustomModel#s3_uri}.
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

