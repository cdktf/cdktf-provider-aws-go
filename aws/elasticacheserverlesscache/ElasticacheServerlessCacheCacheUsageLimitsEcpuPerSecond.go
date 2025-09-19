// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticacheserverlesscache


type ElasticacheServerlessCacheCacheUsageLimitsEcpuPerSecond struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/elasticache_serverless_cache#maximum ElasticacheServerlessCache#maximum}.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/elasticache_serverless_cache#minimum ElasticacheServerlessCache#minimum}.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

