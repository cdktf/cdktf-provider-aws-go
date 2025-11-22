// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallvpcendpointassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallVpcEndpointAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#firewall_arn NetworkfirewallVpcEndpointAssociation#firewall_arn}.
	FirewallArn *string `field:"required" json:"firewallArn" yaml:"firewallArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#vpc_id NetworkfirewallVpcEndpointAssociation#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#description NetworkfirewallVpcEndpointAssociation#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#region NetworkfirewallVpcEndpointAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// subnet_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#subnet_mapping NetworkfirewallVpcEndpointAssociation#subnet_mapping}
	SubnetMapping interface{} `field:"optional" json:"subnetMapping" yaml:"subnetMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#tags NetworkfirewallVpcEndpointAssociation#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.22.1/docs/resources/networkfirewall_vpc_endpoint_association#timeouts NetworkfirewallVpcEndpointAssociation#timeouts}
	Timeouts *NetworkfirewallVpcEndpointAssociationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

