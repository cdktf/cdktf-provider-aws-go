// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerflowdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SagemakerFlowDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#flow_definition_name SagemakerFlowDefinition#flow_definition_name}.
	FlowDefinitionName *string `field:"required" json:"flowDefinitionName" yaml:"flowDefinitionName"`
	// human_loop_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#human_loop_config SagemakerFlowDefinition#human_loop_config}
	HumanLoopConfig *SagemakerFlowDefinitionHumanLoopConfig `field:"required" json:"humanLoopConfig" yaml:"humanLoopConfig"`
	// output_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#output_config SagemakerFlowDefinition#output_config}
	OutputConfig *SagemakerFlowDefinitionOutputConfig `field:"required" json:"outputConfig" yaml:"outputConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#role_arn SagemakerFlowDefinition#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// human_loop_activation_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#human_loop_activation_config SagemakerFlowDefinition#human_loop_activation_config}
	HumanLoopActivationConfig *SagemakerFlowDefinitionHumanLoopActivationConfig `field:"optional" json:"humanLoopActivationConfig" yaml:"humanLoopActivationConfig"`
	// human_loop_request_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#human_loop_request_source SagemakerFlowDefinition#human_loop_request_source}
	HumanLoopRequestSource *SagemakerFlowDefinitionHumanLoopRequestSource `field:"optional" json:"humanLoopRequestSource" yaml:"humanLoopRequestSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#id SagemakerFlowDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#tags SagemakerFlowDefinition#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.0/docs/resources/sagemaker_flow_definition#tags_all SagemakerFlowDefinition#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

