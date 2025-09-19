// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperRoleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/prometheus_scraper#source_role_arn PrometheusScraper#source_role_arn}.
	SourceRoleArn *string `field:"optional" json:"sourceRoleArn" yaml:"sourceRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/prometheus_scraper#target_role_arn PrometheusScraper#target_role_arn}.
	TargetRoleArn *string `field:"optional" json:"targetRoleArn" yaml:"targetRoleArn"`
}

