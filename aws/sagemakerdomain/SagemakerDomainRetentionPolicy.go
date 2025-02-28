// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerdomain


type SagemakerDomainRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.89.0/docs/resources/sagemaker_domain#home_efs_file_system SagemakerDomain#home_efs_file_system}.
	HomeEfsFileSystem *string `field:"optional" json:"homeEfsFileSystem" yaml:"homeEfsFileSystem"`
}

