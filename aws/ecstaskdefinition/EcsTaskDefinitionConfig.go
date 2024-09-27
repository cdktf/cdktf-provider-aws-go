// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecstaskdefinition

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsTaskDefinitionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#container_definitions EcsTaskDefinition#container_definitions}.
	ContainerDefinitions *string `field:"required" json:"containerDefinitions" yaml:"containerDefinitions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#family EcsTaskDefinition#family}.
	Family *string `field:"required" json:"family" yaml:"family"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#cpu EcsTaskDefinition#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// ephemeral_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#ephemeral_storage EcsTaskDefinition#ephemeral_storage}
	EphemeralStorage *EcsTaskDefinitionEphemeralStorage `field:"optional" json:"ephemeralStorage" yaml:"ephemeralStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#execution_role_arn EcsTaskDefinition#execution_role_arn}.
	ExecutionRoleArn *string `field:"optional" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#id EcsTaskDefinition#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// inference_accelerator block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#inference_accelerator EcsTaskDefinition#inference_accelerator}
	InferenceAccelerator interface{} `field:"optional" json:"inferenceAccelerator" yaml:"inferenceAccelerator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#ipc_mode EcsTaskDefinition#ipc_mode}.
	IpcMode *string `field:"optional" json:"ipcMode" yaml:"ipcMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#memory EcsTaskDefinition#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#network_mode EcsTaskDefinition#network_mode}.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#pid_mode EcsTaskDefinition#pid_mode}.
	PidMode *string `field:"optional" json:"pidMode" yaml:"pidMode"`
	// placement_constraints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#placement_constraints EcsTaskDefinition#placement_constraints}
	PlacementConstraints interface{} `field:"optional" json:"placementConstraints" yaml:"placementConstraints"`
	// proxy_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#proxy_configuration EcsTaskDefinition#proxy_configuration}
	ProxyConfiguration *EcsTaskDefinitionProxyConfiguration `field:"optional" json:"proxyConfiguration" yaml:"proxyConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#requires_compatibilities EcsTaskDefinition#requires_compatibilities}.
	RequiresCompatibilities *[]*string `field:"optional" json:"requiresCompatibilities" yaml:"requiresCompatibilities"`
	// runtime_platform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#runtime_platform EcsTaskDefinition#runtime_platform}
	RuntimePlatform *EcsTaskDefinitionRuntimePlatform `field:"optional" json:"runtimePlatform" yaml:"runtimePlatform"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#skip_destroy EcsTaskDefinition#skip_destroy}.
	SkipDestroy interface{} `field:"optional" json:"skipDestroy" yaml:"skipDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#tags EcsTaskDefinition#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#tags_all EcsTaskDefinition#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#task_role_arn EcsTaskDefinition#task_role_arn}.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#track_latest EcsTaskDefinition#track_latest}.
	TrackLatest interface{} `field:"optional" json:"trackLatest" yaml:"trackLatest"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/ecs_task_definition#volume EcsTaskDefinition#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}

