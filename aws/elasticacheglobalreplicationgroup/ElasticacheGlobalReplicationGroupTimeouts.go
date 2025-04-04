// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticacheglobalreplicationgroup


type ElasticacheGlobalReplicationGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/elasticache_global_replication_group#create ElasticacheGlobalReplicationGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/elasticache_global_replication_group#delete ElasticacheGlobalReplicationGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.94.0/docs/resources/elasticache_global_replication_group#update ElasticacheGlobalReplicationGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

