// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation


type CodegurureviewerRepositoryAssociationRepositoryS3Bucket struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/codegurureviewer_repository_association#bucket_name CodegurureviewerRepositoryAssociation#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/codegurureviewer_repository_association#name CodegurureviewerRepositoryAssociation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

