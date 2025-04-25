// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lightsaildistribution


type LightsailDistributionOrigin struct {
	// The name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/lightsail_distribution#name LightsailDistribution#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The AWS Region name of the origin resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/lightsail_distribution#region_name LightsailDistribution#region_name}
	RegionName *string `field:"required" json:"regionName" yaml:"regionName"`
	// The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.96.0/docs/resources/lightsail_distribution#protocol_policy LightsailDistribution#protocol_policy}
	ProtocolPolicy *string `field:"optional" json:"protocolPolicy" yaml:"protocolPolicy"`
}

