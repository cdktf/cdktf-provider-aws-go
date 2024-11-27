// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package finspacekxcluster


type FinspaceKxClusterDatabaseCacheConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/finspace_kx_cluster#cache_type FinspaceKxCluster#cache_type}.
	CacheType *string `field:"required" json:"cacheType" yaml:"cacheType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/finspace_kx_cluster#db_paths FinspaceKxCluster#db_paths}.
	DbPaths *[]*string `field:"optional" json:"dbPaths" yaml:"dbPaths"`
}

