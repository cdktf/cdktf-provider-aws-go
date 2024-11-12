// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpointconfiguration


type SagemakerEndpointConfigurationProductionVariantsRoutingConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/sagemaker_endpoint_configuration#routing_strategy SagemakerEndpointConfiguration#routing_strategy}.
	RoutingStrategy *string `field:"required" json:"routingStrategy" yaml:"routingStrategy"`
}

