// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperSource struct {
	// eks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/prometheus_scraper#eks PrometheusScraper#eks}
	Eks interface{} `field:"optional" json:"eks" yaml:"eks"`
}

