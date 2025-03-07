// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticachereservedcachenode

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticacheReservedCacheNodeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticache_reserved_cache_node#reserved_cache_nodes_offering_id ElasticacheReservedCacheNode#reserved_cache_nodes_offering_id}.
	ReservedCacheNodesOfferingId *string `field:"required" json:"reservedCacheNodesOfferingId" yaml:"reservedCacheNodesOfferingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticache_reserved_cache_node#cache_node_count ElasticacheReservedCacheNode#cache_node_count}.
	CacheNodeCount *float64 `field:"optional" json:"cacheNodeCount" yaml:"cacheNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticache_reserved_cache_node#id ElasticacheReservedCacheNode#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticache_reserved_cache_node#tags ElasticacheReservedCacheNode#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.90.0/docs/resources/elasticache_reserved_cache_node#timeouts ElasticacheReservedCacheNode#timeouts}
	Timeouts *ElasticacheReservedCacheNodeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

