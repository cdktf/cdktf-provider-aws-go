// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigatewayusageplan


type ApiGatewayUsagePlanThrottleSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/api_gateway_usage_plan#burst_limit ApiGatewayUsagePlan#burst_limit}.
	BurstLimit *float64 `field:"optional" json:"burstLimit" yaml:"burstLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/api_gateway_usage_plan#rate_limit ApiGatewayUsagePlan#rate_limit}.
	RateLimit *float64 `field:"optional" json:"rateLimit" yaml:"rateLimit"`
}

