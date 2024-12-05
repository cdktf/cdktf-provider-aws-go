// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyspaceskeyspace


type KeyspacesKeyspaceReplicationSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/keyspaces_keyspace#region_list KeyspacesKeyspace#region_list}.
	RegionList *[]*string `field:"optional" json:"regionList" yaml:"regionList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/keyspaces_keyspace#replication_strategy KeyspacesKeyspace#replication_strategy}.
	ReplicationStrategy *string `field:"optional" json:"replicationStrategy" yaml:"replicationStrategy"`
}

