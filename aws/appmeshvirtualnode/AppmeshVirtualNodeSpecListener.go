// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecListener struct {
	// port_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#port_mapping AppmeshVirtualNode#port_mapping}
	PortMapping *AppmeshVirtualNodeSpecListenerPortMapping `field:"required" json:"portMapping" yaml:"portMapping"`
	// connection_pool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#connection_pool AppmeshVirtualNode#connection_pool}
	ConnectionPool *AppmeshVirtualNodeSpecListenerConnectionPool `field:"optional" json:"connectionPool" yaml:"connectionPool"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#health_check AppmeshVirtualNode#health_check}
	HealthCheck *AppmeshVirtualNodeSpecListenerHealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// outlier_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#outlier_detection AppmeshVirtualNode#outlier_detection}
	OutlierDetection *AppmeshVirtualNodeSpecListenerOutlierDetection `field:"optional" json:"outlierDetection" yaml:"outlierDetection"`
	// timeout block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#timeout AppmeshVirtualNode#timeout}
	Timeout *AppmeshVirtualNodeSpecListenerTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.25.0/docs/resources/appmesh_virtual_node#tls AppmeshVirtualNode#tls}
	Tls *AppmeshVirtualNodeSpecListenerTls `field:"optional" json:"tls" yaml:"tls"`
}

