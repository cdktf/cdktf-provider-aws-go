// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterDatabase struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/finspace_kx_cluster#database_name FinspaceKxCluster#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// cache_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/finspace_kx_cluster#cache_configurations FinspaceKxCluster#cache_configurations}
	CacheConfigurations interface{} `field:"optional" json:"cacheConfigurations" yaml:"cacheConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/finspace_kx_cluster#changeset_id FinspaceKxCluster#changeset_id}.
	ChangesetId *string `field:"optional" json:"changesetId" yaml:"changesetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.42.0/docs/resources/finspace_kx_cluster#dataview_name FinspaceKxCluster#dataview_name}.
	DataviewName *string `field:"optional" json:"dataviewName" yaml:"dataviewName"`
}

