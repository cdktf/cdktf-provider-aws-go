// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2transitgatewaymulticastdomain

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Ec2TransitGatewayMulticastDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#transit_gateway_id Ec2TransitGatewayMulticastDomain#transit_gateway_id}.
	TransitGatewayId *string `field:"required" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#auto_accept_shared_associations Ec2TransitGatewayMulticastDomain#auto_accept_shared_associations}.
	AutoAcceptSharedAssociations *string `field:"optional" json:"autoAcceptSharedAssociations" yaml:"autoAcceptSharedAssociations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#id Ec2TransitGatewayMulticastDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#igmpv2_support Ec2TransitGatewayMulticastDomain#igmpv2_support}.
	Igmpv2Support *string `field:"optional" json:"igmpv2Support" yaml:"igmpv2Support"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#region Ec2TransitGatewayMulticastDomain#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#static_sources_support Ec2TransitGatewayMulticastDomain#static_sources_support}.
	StaticSourcesSupport *string `field:"optional" json:"staticSourcesSupport" yaml:"staticSourcesSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#tags Ec2TransitGatewayMulticastDomain#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#tags_all Ec2TransitGatewayMulticastDomain#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ec2_transit_gateway_multicast_domain#timeouts Ec2TransitGatewayMulticastDomain#timeouts}
	Timeouts *Ec2TransitGatewayMulticastDomainTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

