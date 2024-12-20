// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain


type ElasticsearchDomainAdvancedSecurityOptionsMasterUserOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/elasticsearch_domain#master_user_arn ElasticsearchDomain#master_user_arn}.
	MasterUserArn *string `field:"optional" json:"masterUserArn" yaml:"masterUserArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/elasticsearch_domain#master_user_name ElasticsearchDomain#master_user_name}.
	MasterUserName *string `field:"optional" json:"masterUserName" yaml:"masterUserName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/elasticsearch_domain#master_user_password ElasticsearchDomain#master_user_password}.
	MasterUserPassword *string `field:"optional" json:"masterUserPassword" yaml:"masterUserPassword"`
}

