// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableTtl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.1/docs/resources/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

