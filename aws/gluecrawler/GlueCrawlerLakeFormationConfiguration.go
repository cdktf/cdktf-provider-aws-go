// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gluecrawler


type GlueCrawlerLakeFormationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/glue_crawler#account_id GlueCrawler#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/glue_crawler#use_lake_formation_credentials GlueCrawler#use_lake_formation_credentials}.
	UseLakeFormationCredentials interface{} `field:"optional" json:"useLakeFormationCredentials" yaml:"useLakeFormationCredentials"`
}

