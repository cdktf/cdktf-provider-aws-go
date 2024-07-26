// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation


type CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/codegurureviewer_repository_association#connection_arn CodegurureviewerRepositoryAssociation#connection_arn}.
	ConnectionArn *string `field:"required" json:"connectionArn" yaml:"connectionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/codegurureviewer_repository_association#name CodegurureviewerRepositoryAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.60.0/docs/resources/codegurureviewer_repository_association#owner CodegurureviewerRepositoryAssociation#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
}

