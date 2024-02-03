// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratorcustomroutinglistener

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorCustomRoutingListenerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_custom_routing_listener#accelerator_arn GlobalacceleratorCustomRoutingListener#accelerator_arn}.
	AcceleratorArn *string `field:"required" json:"acceleratorArn" yaml:"acceleratorArn"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_custom_routing_listener#port_range GlobalacceleratorCustomRoutingListener#port_range}
	PortRange interface{} `field:"required" json:"portRange" yaml:"portRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_custom_routing_listener#id GlobalacceleratorCustomRoutingListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_custom_routing_listener#timeouts GlobalacceleratorCustomRoutingListener#timeouts}
	Timeouts *GlobalacceleratorCustomRoutingListenerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

