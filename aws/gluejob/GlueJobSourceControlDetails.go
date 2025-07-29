// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluejob


type GlueJobSourceControlDetails struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#auth_strategy GlueJob#auth_strategy}.
	AuthStrategy *string `field:"optional" json:"authStrategy" yaml:"authStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#auth_token GlueJob#auth_token}.
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#branch GlueJob#branch}.
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#folder GlueJob#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#last_commit_id GlueJob#last_commit_id}.
	LastCommitId *string `field:"optional" json:"lastCommitId" yaml:"lastCommitId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#owner GlueJob#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#provider GlueJob#provider}.
	Provider *string `field:"optional" json:"provider" yaml:"provider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.6.0/docs/resources/glue_job#repository GlueJob#repository}.
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

