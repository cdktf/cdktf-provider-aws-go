// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice


type EcsExpressGatewayServicePrimaryContainerEnvironment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#name EcsExpressGatewayService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ecs_express_gateway_service#value EcsExpressGatewayService#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

