// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package memorydbuser


type MemorydbUserAuthenticationMode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/memorydb_user#type MemorydbUser#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/memorydb_user#passwords MemorydbUser#passwords}.
	Passwords *[]*string `field:"optional" json:"passwords" yaml:"passwords"`
}

