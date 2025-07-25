// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspacesdirectory


type WorkspacesDirectoryCertificateBasedAuthProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/workspaces_directory#certificate_authority_arn WorkspacesDirectory#certificate_authority_arn}.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.5.0/docs/resources/workspaces_directory#status WorkspacesDirectory#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

