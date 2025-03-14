// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualgateway


type AppmeshVirtualGatewaySpecLogging struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.91.0/docs/resources/appmesh_virtual_gateway#access_log AppmeshVirtualGateway#access_log}
	AccessLog *AppmeshVirtualGatewaySpecLoggingAccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
}

