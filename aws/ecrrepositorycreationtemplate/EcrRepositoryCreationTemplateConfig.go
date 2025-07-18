// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecrrepositorycreationtemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcrRepositoryCreationTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#applied_for EcrRepositoryCreationTemplate#applied_for}.
	AppliedFor *[]*string `field:"required" json:"appliedFor" yaml:"appliedFor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#prefix EcrRepositoryCreationTemplate#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#custom_role_arn EcrRepositoryCreationTemplate#custom_role_arn}.
	CustomRoleArn *string `field:"optional" json:"customRoleArn" yaml:"customRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#description EcrRepositoryCreationTemplate#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#encryption_configuration EcrRepositoryCreationTemplate#encryption_configuration}
	EncryptionConfiguration interface{} `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#id EcrRepositoryCreationTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#image_tag_mutability EcrRepositoryCreationTemplate#image_tag_mutability}.
	ImageTagMutability *string `field:"optional" json:"imageTagMutability" yaml:"imageTagMutability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#lifecycle_policy EcrRepositoryCreationTemplate#lifecycle_policy}.
	LifecyclePolicy *string `field:"optional" json:"lifecyclePolicy" yaml:"lifecyclePolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#region EcrRepositoryCreationTemplate#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#repository_policy EcrRepositoryCreationTemplate#repository_policy}.
	RepositoryPolicy *string `field:"optional" json:"repositoryPolicy" yaml:"repositoryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/ecr_repository_creation_template#resource_tags EcrRepositoryCreationTemplate#resource_tags}.
	ResourceTags *map[string]*string `field:"optional" json:"resourceTags" yaml:"resourceTags"`
}

