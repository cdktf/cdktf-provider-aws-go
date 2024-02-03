// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalacceleratoraccelerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GlobalacceleratorAcceleratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#name GlobalacceleratorAccelerator#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// attributes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#attributes GlobalacceleratorAccelerator#attributes}
	Attributes *GlobalacceleratorAcceleratorAttributes `field:"optional" json:"attributes" yaml:"attributes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#enabled GlobalacceleratorAccelerator#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#id GlobalacceleratorAccelerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#ip_addresses GlobalacceleratorAccelerator#ip_addresses}.
	IpAddresses *[]*string `field:"optional" json:"ipAddresses" yaml:"ipAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#ip_address_type GlobalacceleratorAccelerator#ip_address_type}.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#tags GlobalacceleratorAccelerator#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#tags_all GlobalacceleratorAccelerator#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/globalaccelerator_accelerator#timeouts GlobalacceleratorAccelerator#timeouts}
	Timeouts *GlobalacceleratorAcceleratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

