// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package osispipeline

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OsisPipelineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#max_units OsisPipeline#max_units}.
	MaxUnits *float64 `field:"required" json:"maxUnits" yaml:"maxUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#min_units OsisPipeline#min_units}.
	MinUnits *float64 `field:"required" json:"minUnits" yaml:"minUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#pipeline_configuration_body OsisPipeline#pipeline_configuration_body}.
	PipelineConfigurationBody *string `field:"required" json:"pipelineConfigurationBody" yaml:"pipelineConfigurationBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#pipeline_name OsisPipeline#pipeline_name}.
	PipelineName *string `field:"required" json:"pipelineName" yaml:"pipelineName"`
	// buffer_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#buffer_options OsisPipeline#buffer_options}
	BufferOptions interface{} `field:"optional" json:"bufferOptions" yaml:"bufferOptions"`
	// encryption_at_rest_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#encryption_at_rest_options OsisPipeline#encryption_at_rest_options}
	EncryptionAtRestOptions interface{} `field:"optional" json:"encryptionAtRestOptions" yaml:"encryptionAtRestOptions"`
	// log_publishing_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#log_publishing_options OsisPipeline#log_publishing_options}
	LogPublishingOptions interface{} `field:"optional" json:"logPublishingOptions" yaml:"logPublishingOptions"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#region OsisPipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#tags OsisPipeline#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#timeouts OsisPipeline#timeouts}
	Timeouts *OsisPipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/osis_pipeline#vpc_options OsisPipeline#vpc_options}
	VpcOptions interface{} `field:"optional" json:"vpcOptions" yaml:"vpcOptions"`
}

