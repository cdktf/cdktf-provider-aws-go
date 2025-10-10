// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperDestinationAmp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/prometheus_scraper#workspace_arn PrometheusScraper#workspace_arn}.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
}

