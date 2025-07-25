// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectSecondarySources struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#source_identifier CodebuildProject#source_identifier}.
	SourceIdentifier *string `field:"required" json:"sourceIdentifier" yaml:"sourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#type CodebuildProject#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#auth CodebuildProject#auth}
	Auth *CodebuildProjectSecondarySourcesAuth `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#buildspec CodebuildProject#buildspec}.
	Buildspec *string `field:"optional" json:"buildspec" yaml:"buildspec"`
	// build_status_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#build_status_config CodebuildProject#build_status_config}
	BuildStatusConfig *CodebuildProjectSecondarySourcesBuildStatusConfig `field:"optional" json:"buildStatusConfig" yaml:"buildStatusConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#git_clone_depth CodebuildProject#git_clone_depth}.
	GitCloneDepth *float64 `field:"optional" json:"gitCloneDepth" yaml:"gitCloneDepth"`
	// git_submodules_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#git_submodules_config CodebuildProject#git_submodules_config}
	GitSubmodulesConfig *CodebuildProjectSecondarySourcesGitSubmodulesConfig `field:"optional" json:"gitSubmodulesConfig" yaml:"gitSubmodulesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#insecure_ssl CodebuildProject#insecure_ssl}.
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#location CodebuildProject#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/codebuild_project#report_build_status CodebuildProject#report_build_status}.
	ReportBuildStatus interface{} `field:"optional" json:"reportBuildStatus" yaml:"reportBuildStatus"`
}

