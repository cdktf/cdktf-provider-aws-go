// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain


type ElasticsearchDomainAutoTuneOptionsMaintenanceScheduleDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elasticsearch_domain#unit ElasticsearchDomain#unit}.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.83.1/docs/resources/elasticsearch_domain#value ElasticsearchDomain#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

