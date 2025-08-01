// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkfirewallfirewallpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkfirewallFirewallPolicyConfig struct {
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
	// firewall_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#firewall_policy NetworkfirewallFirewallPolicy#firewall_policy}
	FirewallPolicy *NetworkfirewallFirewallPolicyFirewallPolicy `field:"required" json:"firewallPolicy" yaml:"firewallPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#name NetworkfirewallFirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#description NetworkfirewallFirewallPolicy#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#encryption_configuration NetworkfirewallFirewallPolicy#encryption_configuration}
	EncryptionConfiguration *NetworkfirewallFirewallPolicyEncryptionConfiguration `field:"optional" json:"encryptionConfiguration" yaml:"encryptionConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#id NetworkfirewallFirewallPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#region NetworkfirewallFirewallPolicy#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#tags NetworkfirewallFirewallPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/networkfirewall_firewall_policy#tags_all NetworkfirewallFirewallPolicy#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

