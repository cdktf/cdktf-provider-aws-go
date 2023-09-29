// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspaceskeyspace


type KeyspacesKeyspaceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/keyspaces_keyspace#create KeyspacesKeyspace#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.19.0/docs/resources/keyspaces_keyspace#delete KeyspacesKeyspace#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

