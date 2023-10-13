// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydbuser


type MemorydbUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/memorydb_user#passwords MemorydbUser#passwords}.
	Passwords *[]*string `field:"required" json:"passwords" yaml:"passwords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/memorydb_user#type MemorydbUser#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

