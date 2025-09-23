// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain


type ElasticsearchDomainAdvancedSecurityOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticsearch_domain#internal_user_database_enabled ElasticsearchDomain#internal_user_database_enabled}.
	InternalUserDatabaseEnabled interface{} `field:"optional" json:"internalUserDatabaseEnabled" yaml:"internalUserDatabaseEnabled"`
	// master_user_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.1/docs/resources/elasticsearch_domain#master_user_options ElasticsearchDomain#master_user_options}
	MasterUserOptions *ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions `field:"optional" json:"masterUserOptions" yaml:"masterUserOptions"`
}

