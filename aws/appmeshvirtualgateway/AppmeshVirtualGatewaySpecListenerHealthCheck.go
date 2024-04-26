// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListenerHealthCheck struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#healthy_threshold AppmeshVirtualGateway#healthy_threshold}.
	HealthyThreshold *float64 `field:"required" json:"healthyThreshold" yaml:"healthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#interval_millis AppmeshVirtualGateway#interval_millis}.
	IntervalMillis *float64 `field:"required" json:"intervalMillis" yaml:"intervalMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#protocol AppmeshVirtualGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#timeout_millis AppmeshVirtualGateway#timeout_millis}.
	TimeoutMillis *float64 `field:"required" json:"timeoutMillis" yaml:"timeoutMillis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#unhealthy_threshold AppmeshVirtualGateway#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#path AppmeshVirtualGateway#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.47.0/docs/resources/appmesh_virtual_gateway#port AppmeshVirtualGateway#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

