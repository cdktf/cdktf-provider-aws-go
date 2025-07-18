// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableTtl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

