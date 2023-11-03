// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dbclustersnapshot


type DbClusterSnapshotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.24.0/docs/resources/db_cluster_snapshot#create DbClusterSnapshot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
}

