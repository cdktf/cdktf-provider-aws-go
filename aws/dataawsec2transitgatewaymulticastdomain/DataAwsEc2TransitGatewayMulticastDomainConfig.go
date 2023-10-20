// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsec2transitgatewaymulticastdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsEc2TransitGatewayMulticastDomainConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/ec2_transit_gateway_multicast_domain#filter DataAwsEc2TransitGatewayMulticastDomain#filter}
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/ec2_transit_gateway_multicast_domain#id DataAwsEc2TransitGatewayMulticastDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/ec2_transit_gateway_multicast_domain#tags DataAwsEc2TransitGatewayMulticastDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/ec2_transit_gateway_multicast_domain#timeouts DataAwsEc2TransitGatewayMulticastDomain#timeouts}
	Timeouts *DataAwsEc2TransitGatewayMulticastDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.22.0/docs/data-sources/ec2_transit_gateway_multicast_domain#transit_gateway_multicast_domain_id DataAwsEc2TransitGatewayMulticastDomain#transit_gateway_multicast_domain_id}.
	TransitGatewayMulticastDomainId *string `field:"optional" json:"transitGatewayMulticastDomainId" yaml:"transitGatewayMulticastDomainId"`
}

