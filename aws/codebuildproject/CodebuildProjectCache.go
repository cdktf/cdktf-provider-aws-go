// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectCache struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/codebuild_project#location CodebuildProject#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/codebuild_project#modes CodebuildProject#modes}.
	Modes *[]*string `field:"optional" json:"modes" yaml:"modes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

