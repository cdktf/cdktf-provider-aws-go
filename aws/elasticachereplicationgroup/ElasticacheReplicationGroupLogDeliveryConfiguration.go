// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticachereplicationgroup


type ElasticacheReplicationGroupLogDeliveryConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/elasticache_replication_group#destination ElasticacheReplicationGroup#destination}.
	Destination *string `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/elasticache_replication_group#destination_type ElasticacheReplicationGroup#destination_type}.
	DestinationType *string `field:"required" json:"destinationType" yaml:"destinationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/elasticache_replication_group#log_format ElasticacheReplicationGroup#log_format}.
	LogFormat *string `field:"required" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.11.0/docs/resources/elasticache_replication_group#log_type ElasticacheReplicationGroup#log_type}.
	LogType *string `field:"required" json:"logType" yaml:"logType"`
}

