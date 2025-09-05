// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewalltransitgatewayattachmentaccepter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallFirewallTransitGatewayAttachmentAccepterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/networkfirewall_firewall_transit_gateway_attachment_accepter#transit_gateway_attachment_id NetworkfirewallFirewallTransitGatewayAttachmentAccepter#transit_gateway_attachment_id}.
	TransitGatewayAttachmentId *string `field:"required" json:"transitGatewayAttachmentId" yaml:"transitGatewayAttachmentId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/networkfirewall_firewall_transit_gateway_attachment_accepter#region NetworkfirewallFirewallTransitGatewayAttachmentAccepter#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/networkfirewall_firewall_transit_gateway_attachment_accepter#timeouts NetworkfirewallFirewallTransitGatewayAttachmentAccepter#timeouts}
	Timeouts *NetworkfirewallFirewallTransitGatewayAttachmentAccepterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

