// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bedrockcustommodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BedrockCustomModelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#base_model_identifier BedrockCustomModel#base_model_identifier}.
	BaseModelIdentifier *string `field:"required" json:"baseModelIdentifier" yaml:"baseModelIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#custom_model_name BedrockCustomModel#custom_model_name}.
	CustomModelName *string `field:"required" json:"customModelName" yaml:"customModelName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#hyperparameters BedrockCustomModel#hyperparameters}.
	Hyperparameters *map[string]*string `field:"required" json:"hyperparameters" yaml:"hyperparameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#job_name BedrockCustomModel#job_name}.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#role_arn BedrockCustomModel#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#customization_type BedrockCustomModel#customization_type}.
	CustomizationType *string `field:"optional" json:"customizationType" yaml:"customizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#custom_model_kms_key_id BedrockCustomModel#custom_model_kms_key_id}.
	CustomModelKmsKeyId *string `field:"optional" json:"customModelKmsKeyId" yaml:"customModelKmsKeyId"`
	// output_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#output_data_config BedrockCustomModel#output_data_config}
	OutputDataConfig interface{} `field:"optional" json:"outputDataConfig" yaml:"outputDataConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#tags BedrockCustomModel#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#timeouts BedrockCustomModel#timeouts}
	Timeouts *BedrockCustomModelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// training_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#training_data_config BedrockCustomModel#training_data_config}
	TrainingDataConfig interface{} `field:"optional" json:"trainingDataConfig" yaml:"trainingDataConfig"`
	// validation_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#validation_data_config BedrockCustomModel#validation_data_config}
	ValidationDataConfig interface{} `field:"optional" json:"validationDataConfig" yaml:"validationDataConfig"`
	// vpc_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.46.0/docs/resources/bedrock_custom_model#vpc_config BedrockCustomModel#vpc_config}
	VpcConfig interface{} `field:"optional" json:"vpcConfig" yaml:"vpcConfig"`
}

