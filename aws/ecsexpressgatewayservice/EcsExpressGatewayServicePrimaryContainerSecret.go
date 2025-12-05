// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ecsexpressgatewayservice


type EcsExpressGatewayServicePrimaryContainerSecret struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_express_gateway_service#name EcsExpressGatewayService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/ecs_express_gateway_service#value_from EcsExpressGatewayService#value_from}.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

