// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EcsExpressGatewayServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#execution_role_arn EcsExpressGatewayService#execution_role_arn}.
	ExecutionRoleArn *string `field:"required" json:"executionRoleArn" yaml:"executionRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#infrastructure_role_arn EcsExpressGatewayService#infrastructure_role_arn}.
	InfrastructureRoleArn *string `field:"required" json:"infrastructureRoleArn" yaml:"infrastructureRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#cluster EcsExpressGatewayService#cluster}.
	Cluster *string `field:"optional" json:"cluster" yaml:"cluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#cpu EcsExpressGatewayService#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#health_check_path EcsExpressGatewayService#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#memory EcsExpressGatewayService#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#network_configuration EcsExpressGatewayService#network_configuration}.
	NetworkConfiguration interface{} `field:"optional" json:"networkConfiguration" yaml:"networkConfiguration"`
	// primary_container block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#primary_container EcsExpressGatewayService#primary_container}
	PrimaryContainer interface{} `field:"optional" json:"primaryContainer" yaml:"primaryContainer"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#region EcsExpressGatewayService#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#scaling_target EcsExpressGatewayService#scaling_target}.
	ScalingTarget interface{} `field:"optional" json:"scalingTarget" yaml:"scalingTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#service_name EcsExpressGatewayService#service_name}.
	ServiceName *string `field:"optional" json:"serviceName" yaml:"serviceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#tags EcsExpressGatewayService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#task_role_arn EcsExpressGatewayService#task_role_arn}.
	TaskRoleArn *string `field:"optional" json:"taskRoleArn" yaml:"taskRoleArn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#timeouts EcsExpressGatewayService#timeouts}
	Timeouts *EcsExpressGatewayServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#wait_for_steady_state EcsExpressGatewayService#wait_for_steady_state}.
	WaitForSteadyState interface{} `field:"optional" json:"waitForSteadyState" yaml:"waitForSteadyState"`
}

