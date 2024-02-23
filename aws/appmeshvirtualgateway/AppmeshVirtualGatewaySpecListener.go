// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecListener struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/appmesh_virtual_gateway#port_mapping AppmeshVirtualGateway#port_mapping}
	PortMapping *AppmeshVirtualGatewaySpecListenerPortMapping `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/appmesh_virtual_gateway#connection_pool AppmeshVirtualGateway#connection_pool}
	ConnectionPool *AppmeshVirtualGatewaySpecListenerConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/appmesh_virtual_gateway#health_check AppmeshVirtualGateway#health_check}
	HealthCheck *AppmeshVirtualGatewaySpecListenerHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.38.0/docs/resources/appmesh_virtual_gateway#tls AppmeshVirtualGateway#tls}
	Tls *AppmeshVirtualGatewaySpecListenerTls `field:"optional" json:"tls" yaml:"tls"`
}

