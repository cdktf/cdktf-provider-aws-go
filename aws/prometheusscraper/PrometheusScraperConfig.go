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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#scrape_configuration PrometheusScraper#scrape_configuration}.
	ScrapeConfiguration *string `field:"required" json:"scrapeConfiguration" yaml:"scrapeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#alias PrometheusScraper#alias}.
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#destination PrometheusScraper#destination}
	Destination interface{} `field:"optional" json:"destination" yaml:"destination"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#region PrometheusScraper#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// role_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#role_configuration PrometheusScraper#role_configuration}
	RoleConfiguration interface{} `field:"optional" json:"roleConfiguration" yaml:"roleConfiguration"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#source PrometheusScraper#source}
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#tags PrometheusScraper#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/prometheus_scraper#timeouts PrometheusScraper#timeouts}
	Timeouts *PrometheusScraperTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

