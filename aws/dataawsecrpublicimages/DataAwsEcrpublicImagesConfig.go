// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsecrpublicimages

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEcrpublicImagesConfig struct {
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
	// Name of the public repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/data-sources/ecrpublic_images#repository_name DataAwsEcrpublicImages#repository_name}
	RepositoryName *string `field:"required" json:"repositoryName" yaml:"repositoryName"`
	// image_ids block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/data-sources/ecrpublic_images#image_ids DataAwsEcrpublicImages#image_ids}
	ImageIds interface{} `field:"optional" json:"imageIds" yaml:"imageIds"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/data-sources/ecrpublic_images#region DataAwsEcrpublicImages#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// AWS account ID associated with the public registry that contains the repository.
	//
	// If not specified, the default public registry is assumed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/data-sources/ecrpublic_images#registry_id DataAwsEcrpublicImages#registry_id}
	RegistryId *string `field:"optional" json:"registryId" yaml:"registryId"`
}

