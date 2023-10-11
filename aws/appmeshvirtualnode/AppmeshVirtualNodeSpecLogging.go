// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecLogging struct {
	// access_log block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/appmesh_virtual_node#access_log AppmeshVirtualNode#access_log}
	AccessLog *AppmeshVirtualNodeSpecLoggingAccessLog `field:"optional" json:"accessLog" yaml:"accessLog"`
}

