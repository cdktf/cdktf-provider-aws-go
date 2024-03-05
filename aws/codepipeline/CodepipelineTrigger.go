// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineTrigger struct {
	// git_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/codepipeline#git_configuration Codepipeline#git_configuration}
	GitConfiguration *CodepipelineTriggerGitConfiguration `field:"required" json:"gitConfiguration" yaml:"gitConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.39.1/docs/resources/codepipeline#provider_type Codepipeline#provider_type}.
	ProviderType *string `field:"required" json:"providerType" yaml:"providerType"`
}

