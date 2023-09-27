// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codecatalystdevenvironment


type CodecatalystDevEnvironmentRepositories struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/codecatalyst_dev_environment#repository_name CodecatalystDevEnvironment#repository_name}.
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.18.1/docs/resources/codecatalyst_dev_environment#branch_name CodecatalystDevEnvironment#branch_name}.
	BranchName *string `field:"optional" json:"branchName" yaml:"branchName"`
}

