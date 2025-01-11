// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastictranscoderPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#input_bucket ElastictranscoderPipeline#input_bucket}.
	InputBucket *string `field:"required" json:"inputBucket" yaml:"inputBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#role ElastictranscoderPipeline#role}.
	Role *string `field:"required" json:"role" yaml:"role"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#aws_kms_key_arn ElastictranscoderPipeline#aws_kms_key_arn}.
	AwsKmsKeyArn *string `field:"optional" json:"awsKmsKeyArn" yaml:"awsKmsKeyArn"`
	// content_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#content_config ElastictranscoderPipeline#content_config}
	ContentConfig *ElastictranscoderPipelineContentConfig `field:"optional" json:"contentConfig" yaml:"contentConfig"`
	// content_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#content_config_permissions ElastictranscoderPipeline#content_config_permissions}
	ContentConfigPermissions interface{} `field:"optional" json:"contentConfigPermissions" yaml:"contentConfigPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#id ElastictranscoderPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#name ElastictranscoderPipeline#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#notifications ElastictranscoderPipeline#notifications}
	Notifications *ElastictranscoderPipelineNotifications `field:"optional" json:"notifications" yaml:"notifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#output_bucket ElastictranscoderPipeline#output_bucket}.
	OutputBucket *string `field:"optional" json:"outputBucket" yaml:"outputBucket"`
	// thumbnail_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#thumbnail_config ElastictranscoderPipeline#thumbnail_config}
	ThumbnailConfig *ElastictranscoderPipelineThumbnailConfig `field:"optional" json:"thumbnailConfig" yaml:"thumbnailConfig"`
	// thumbnail_config_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elastictranscoder_pipeline#thumbnail_config_permissions ElastictranscoderPipeline#thumbnail_config_permissions}
	ThumbnailConfigPermissions interface{} `field:"optional" json:"thumbnailConfigPermissions" yaml:"thumbnailConfigPermissions"`
}

