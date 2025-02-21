// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockinferenceprofile


type BedrockInferenceProfileModelSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.88.0/docs/resources/bedrock_inference_profile#copy_from BedrockInferenceProfile#copy_from}.
	CopyFrom *string `field:"required" json:"copyFrom" yaml:"copyFrom"`
}

