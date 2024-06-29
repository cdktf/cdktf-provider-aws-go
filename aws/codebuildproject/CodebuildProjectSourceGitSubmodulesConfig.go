// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codebuildproject


type CodebuildProjectSourceGitSubmodulesConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/codebuild_project#fetch_submodules CodebuildProject#fetch_submodules}.
	FetchSubmodules interface{} `field:"required" json:"fetchSubmodules" yaml:"fetchSubmodules"`
}

