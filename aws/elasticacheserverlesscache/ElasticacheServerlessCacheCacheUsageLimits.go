// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimits struct {
	// data_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/elasticache_serverless_cache#data_storage ElasticacheServerlessCache#data_storage}
	DataStorage interface{} `field:"optional" json:"dataStorage" yaml:"dataStorage"`
	// ecpu_per_second block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/elasticache_serverless_cache#ecpu_per_second ElasticacheServerlessCache#ecpu_per_second}
	EcpuPerSecond interface{} `field:"optional" json:"ecpuPerSecond" yaml:"ecpuPerSecond"`
}

