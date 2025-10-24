// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockagentcorebrowser


type BedrockagentcoreBrowserRecording struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/bedrockagentcore_browser#enabled BedrockagentcoreBrowser#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// s3_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.18.0/docs/resources/bedrockagentcore_browser#s3_location BedrockagentcoreBrowser#s3_location}
	S3Location interface{} `field:"optional" json:"s3Location" yaml:"s3Location"`
}

