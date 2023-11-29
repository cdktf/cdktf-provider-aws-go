// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation


type CodegurureviewerRepositoryAssociationRepository struct {
	// bitbucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/codegurureviewer_repository_association#bitbucket CodegurureviewerRepositoryAssociation#bitbucket}
	Bitbucket *CodegurureviewerRepositoryAssociationRepositoryBitbucket `field:"optional" json:"bitbucket" yaml:"bitbucket"`
	// codecommit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/codegurureviewer_repository_association#codecommit CodegurureviewerRepositoryAssociation#codecommit}
	Codecommit *CodegurureviewerRepositoryAssociationRepositoryCodecommit `field:"optional" json:"codecommit" yaml:"codecommit"`
	// github_enterprise_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/codegurureviewer_repository_association#github_enterprise_server CodegurureviewerRepositoryAssociation#github_enterprise_server}
	GithubEnterpriseServer *CodegurureviewerRepositoryAssociationRepositoryGithubEnterpriseServer `field:"optional" json:"githubEnterpriseServer" yaml:"githubEnterpriseServer"`
	// s3_bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/codegurureviewer_repository_association#s3_bucket CodegurureviewerRepositoryAssociation#s3_bucket}
	S3Bucket *CodegurureviewerRepositoryAssociationRepositoryS3Bucket `field:"optional" json:"s3Bucket" yaml:"s3Bucket"`
}

