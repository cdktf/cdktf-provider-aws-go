// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcustomroutingaccelerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorCustomRoutingAcceleratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#name GlobalacceleratorCustomRoutingAccelerator#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#attributes GlobalacceleratorCustomRoutingAccelerator#attributes}
	Attributes *GlobalacceleratorCustomRoutingAcceleratorAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#enabled GlobalacceleratorCustomRoutingAccelerator#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#id GlobalacceleratorCustomRoutingAccelerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#ip_addresses GlobalacceleratorCustomRoutingAccelerator#ip_addresses}.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#ip_address_type GlobalacceleratorCustomRoutingAccelerator#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#tags GlobalacceleratorCustomRoutingAccelerator#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#tags_all GlobalacceleratorCustomRoutingAccelerator#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/globalaccelerator_custom_routing_accelerator#timeouts GlobalacceleratorCustomRoutingAccelerator#timeouts}
	Timeouts *GlobalacceleratorCustomRoutingAcceleratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

