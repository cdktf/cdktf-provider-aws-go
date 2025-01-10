// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerhub

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerHubConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#hub_description SagemakerHub#hub_description}.
	HubDescription *string `field:"required" json:"hubDescription" yaml:"hubDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#hub_name SagemakerHub#hub_name}.
	HubName *string `field:"required" json:"hubName" yaml:"hubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#hub_display_name SagemakerHub#hub_display_name}.
	HubDisplayName *string `field:"optional" json:"hubDisplayName" yaml:"hubDisplayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#hub_search_keywords SagemakerHub#hub_search_keywords}.
	HubSearchKeywords *[]*string `field:"optional" json:"hubSearchKeywords" yaml:"hubSearchKeywords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#id SagemakerHub#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// s3_storage_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#s3_storage_config SagemakerHub#s3_storage_config}
	S3StorageConfig *SagemakerHubS3StorageConfig `field:"optional" json:"s3StorageConfig" yaml:"s3StorageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#tags SagemakerHub#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.0/docs/resources/sagemaker_hub#tags_all SagemakerHub#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

