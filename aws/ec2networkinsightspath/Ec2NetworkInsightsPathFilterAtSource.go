// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightspath


type Ec2NetworkInsightsPathFilterAtSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#destination_address Ec2NetworkInsightsPath#destination_address}.
	DestinationAddress *string `field:"optional" json:"destinationAddress" yaml:"destinationAddress"`
	// destination_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#destination_port_range Ec2NetworkInsightsPath#destination_port_range}
	DestinationPortRange *Ec2NetworkInsightsPathFilterAtSourceDestinationPortRange `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#source_address Ec2NetworkInsightsPath#source_address}.
	SourceAddress *string `field:"optional" json:"sourceAddress" yaml:"sourceAddress"`
	// source_port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#source_port_range Ec2NetworkInsightsPath#source_port_range}
	SourcePortRange *Ec2NetworkInsightsPathFilterAtSourceSourcePortRange `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
}

