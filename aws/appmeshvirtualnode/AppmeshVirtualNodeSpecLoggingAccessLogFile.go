// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appmeshvirtualnode


type AppmeshVirtualNodeSpecLoggingAccessLogFile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/appmesh_virtual_node#path AppmeshVirtualNode#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.19.0/docs/resources/appmesh_virtual_node#format AppmeshVirtualNode#format}
	Format *AppmeshVirtualNodeSpecLoggingAccessLogFileFormat `field:"optional" json:"format" yaml:"format"`
}

