// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftcluster


type RedshiftClusterSnapshotCopy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/redshift_cluster#destination_region RedshiftCluster#destination_region}.
	DestinationRegion *string `field:"required" json:"destinationRegion" yaml:"destinationRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/redshift_cluster#grant_name RedshiftCluster#grant_name}.
	GrantName *string `field:"optional" json:"grantName" yaml:"grantName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.15.0/docs/resources/redshift_cluster#retention_period RedshiftCluster#retention_period}.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
}

