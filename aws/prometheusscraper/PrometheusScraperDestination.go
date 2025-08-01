// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper


type PrometheusScraperDestination struct {
	// amp block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/prometheus_scraper#amp PrometheusScraper#amp}
	Amp interface{} `field:"optional" json:"amp" yaml:"amp"`
}

