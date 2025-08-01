// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectFileSystemLocations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codebuild_project#identifier CodebuildProject#identifier}.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codebuild_project#location CodebuildProject#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codebuild_project#mount_options CodebuildProject#mount_options}.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codebuild_project#mount_point CodebuildProject#mount_point}.
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

