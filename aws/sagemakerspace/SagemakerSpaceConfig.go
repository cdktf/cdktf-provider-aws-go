// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerspace

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerSpaceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#domain_id SagemakerSpace#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#space_name SagemakerSpace#space_name}.
	SpaceName *string `field:"required" json:"spaceName" yaml:"spaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#id SagemakerSpace#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ownership_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#ownership_settings SagemakerSpace#ownership_settings}
	OwnershipSettings *SagemakerSpaceOwnershipSettings `field:"optional" json:"ownershipSettings" yaml:"ownershipSettings"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#region SagemakerSpace#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#space_display_name SagemakerSpace#space_display_name}.
	SpaceDisplayName *string `field:"optional" json:"spaceDisplayName" yaml:"spaceDisplayName"`
	// space_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#space_settings SagemakerSpace#space_settings}
	SpaceSettings *SagemakerSpaceSpaceSettings `field:"optional" json:"spaceSettings" yaml:"spaceSettings"`
	// space_sharing_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#space_sharing_settings SagemakerSpace#space_sharing_settings}
	SpaceSharingSettings *SagemakerSpaceSpaceSharingSettings `field:"optional" json:"spaceSharingSettings" yaml:"spaceSharingSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#tags SagemakerSpace#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.9.0/docs/resources/sagemaker_space#tags_all SagemakerSpace#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

