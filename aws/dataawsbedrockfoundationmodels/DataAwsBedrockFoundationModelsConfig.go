// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsbedrockfoundationmodels

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsBedrockFoundationModelsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/data-sources/bedrock_foundation_models#by_customization_type DataAwsBedrockFoundationModels#by_customization_type}.
	ByCustomizationType *string `field:"optional" json:"byCustomizationType" yaml:"byCustomizationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/data-sources/bedrock_foundation_models#by_inference_type DataAwsBedrockFoundationModels#by_inference_type}.
	ByInferenceType *string `field:"optional" json:"byInferenceType" yaml:"byInferenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/data-sources/bedrock_foundation_models#by_output_modality DataAwsBedrockFoundationModels#by_output_modality}.
	ByOutputModality *string `field:"optional" json:"byOutputModality" yaml:"byOutputModality"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/data-sources/bedrock_foundation_models#by_provider DataAwsBedrockFoundationModels#by_provider}.
	ByProvider *string `field:"optional" json:"byProvider" yaml:"byProvider"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/data-sources/bedrock_foundation_models#region DataAwsBedrockFoundationModels#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

