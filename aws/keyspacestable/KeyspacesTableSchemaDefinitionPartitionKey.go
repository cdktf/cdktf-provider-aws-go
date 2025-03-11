// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspacestable


type KeyspacesTableSchemaDefinitionPartitionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.1/docs/resources/keyspaces_table#name KeyspacesTable#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

