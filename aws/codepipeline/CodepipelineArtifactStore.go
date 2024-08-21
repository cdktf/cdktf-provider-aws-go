// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipeline


type CodepipelineArtifactStore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codepipeline#location Codepipeline#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codepipeline#type Codepipeline#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// encryption_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codepipeline#encryption_key Codepipeline#encryption_key}
	EncryptionKey *CodepipelineArtifactStoreEncryptionKey `field:"optional" json:"encryptionKey" yaml:"encryptionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/codepipeline#region Codepipeline#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

