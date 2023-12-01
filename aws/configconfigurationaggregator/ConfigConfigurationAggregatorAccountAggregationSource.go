// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigurationaggregator


type ConfigConfigurationAggregatorAccountAggregationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/config_configuration_aggregator#account_ids ConfigConfigurationAggregator#account_ids}.
	AccountIds *[]*string `field:"required" json:"accountIds" yaml:"accountIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/config_configuration_aggregator#all_regions ConfigConfigurationAggregator#all_regions}.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.29.0/docs/resources/config_configuration_aggregator#regions ConfigConfigurationAggregator#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

