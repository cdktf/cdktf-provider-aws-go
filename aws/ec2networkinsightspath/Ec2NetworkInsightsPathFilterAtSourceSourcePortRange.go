// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2networkinsightspath


type Ec2NetworkInsightsPathFilterAtSourceSourcePortRange struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#from_port Ec2NetworkInsightsPath#from_port}.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/ec2_network_insights_path#to_port Ec2NetworkInsightsPath#to_port}.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

