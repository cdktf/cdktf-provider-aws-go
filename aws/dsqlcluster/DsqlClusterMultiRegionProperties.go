// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dsqlcluster


type DsqlClusterMultiRegionProperties struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dsql_cluster#clusters DsqlCluster#clusters}.
	Clusters *[]*string `field:"optional" json:"clusters" yaml:"clusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/dsql_cluster#witness_region DsqlCluster#witness_region}.
	WitnessRegion *string `field:"optional" json:"witnessRegion" yaml:"witnessRegion"`
}

