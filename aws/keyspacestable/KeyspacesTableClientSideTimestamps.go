// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableClientSideTimestamps struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

