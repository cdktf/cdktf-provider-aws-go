// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdbcluster


type DocdbClusterRestoreToPointInTime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/docdb_cluster#source_cluster_identifier DocdbCluster#source_cluster_identifier}.
	SourceClusterIdentifier *string `field:"required" json:"sourceClusterIdentifier" yaml:"sourceClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/docdb_cluster#restore_to_time DocdbCluster#restore_to_time}.
	RestoreToTime *string `field:"optional" json:"restoreToTime" yaml:"restoreToTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/docdb_cluster#restore_type DocdbCluster#restore_type}.
	RestoreType *string `field:"optional" json:"restoreType" yaml:"restoreType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.79.0/docs/resources/docdb_cluster#use_latest_restorable_time DocdbCluster#use_latest_restorable_time}.
	UseLatestRestorableTime interface{} `field:"optional" json:"useLatestRestorableTime" yaml:"useLatestRestorableTime"`
}

