// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codepipelinecustomactiontype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodepipelineCustomActionTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#category CodepipelineCustomActionType#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// input_artifact_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#input_artifact_details CodepipelineCustomActionType#input_artifact_details}
	InputArtifactDetails *CodepipelineCustomActionTypeInputArtifactDetails `field:"required" json:"inputArtifactDetails" yaml:"inputArtifactDetails"`
	// output_artifact_details block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#output_artifact_details CodepipelineCustomActionType#output_artifact_details}
	OutputArtifactDetails *CodepipelineCustomActionTypeOutputArtifactDetails `field:"required" json:"outputArtifactDetails" yaml:"outputArtifactDetails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#provider_name CodepipelineCustomActionType#provider_name}.
	ProviderName *string `field:"required" json:"providerName" yaml:"providerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#version CodepipelineCustomActionType#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// configuration_property block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#configuration_property CodepipelineCustomActionType#configuration_property}
	ConfigurationProperty interface{} `field:"optional" json:"configurationProperty" yaml:"configurationProperty"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#id CodepipelineCustomActionType#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#region CodepipelineCustomActionType#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#settings CodepipelineCustomActionType#settings}
	Settings *CodepipelineCustomActionTypeSettings `field:"optional" json:"settings" yaml:"settings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#tags CodepipelineCustomActionType#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/codepipeline_custom_action_type#tags_all CodepipelineCustomActionType#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

