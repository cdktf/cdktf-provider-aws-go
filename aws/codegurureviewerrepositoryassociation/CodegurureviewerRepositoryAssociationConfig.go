// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codegurureviewerrepositoryassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodegurureviewerRepositoryAssociationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#repository CodegurureviewerRepositoryAssociation#repository}
	Repository *CodegurureviewerRepositoryAssociationRepository `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#id CodegurureviewerRepositoryAssociation#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// kms_key_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#kms_key_details CodegurureviewerRepositoryAssociation#kms_key_details}
	KmsKeyDetails *CodegurureviewerRepositoryAssociationKmsKeyDetails `field:"optional" json:"kmsKeyDetails" yaml:"kmsKeyDetails"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#region CodegurureviewerRepositoryAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#tags CodegurureviewerRepositoryAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#tags_all CodegurureviewerRepositoryAssociation#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/codegurureviewer_repository_association#timeouts CodegurureviewerRepositoryAssociation#timeouts}
	Timeouts *CodegurureviewerRepositoryAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

