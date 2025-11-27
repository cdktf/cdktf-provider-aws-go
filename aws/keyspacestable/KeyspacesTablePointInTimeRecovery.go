// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTablePointInTimeRecovery struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/keyspaces_table#status KeyspacesTable#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

