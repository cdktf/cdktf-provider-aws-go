// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elasticsearchdomain


type ElasticsearchDomainCognitoOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/elasticsearch_domain#identity_pool_id ElasticsearchDomain#identity_pool_id}.
	IdentityPoolId *string `field:"required" json:"identityPoolId" yaml:"identityPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/elasticsearch_domain#role_arn ElasticsearchDomain#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/elasticsearch_domain#user_pool_id ElasticsearchDomain#user_pool_id}.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/elasticsearch_domain#enabled ElasticsearchDomain#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

