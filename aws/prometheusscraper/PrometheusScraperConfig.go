// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package prometheusscraper

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrometheusScraperConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#scrape_configuration PrometheusScraper#scrape_configuration}.
	ScrapeConfiguration *string `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#alias PrometheusScraper#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#destination PrometheusScraper#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#source PrometheusScraper#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#tags PrometheusScraper#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/prometheus_scraper#timeouts PrometheusScraper#timeouts}
	Timeouts *PrometheusScraperTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

