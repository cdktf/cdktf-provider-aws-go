// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package configconfigurationaggregator


type ConfigConfigurationAggregatorOrganizationAggregationSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/config_configuration_aggregator#role_arn ConfigConfigurationAggregator#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/config_configuration_aggregator#all_regions ConfigConfigurationAggregator#all_regions}.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/resources/config_configuration_aggregator#regions ConfigConfigurationAggregator#regions}.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
}

