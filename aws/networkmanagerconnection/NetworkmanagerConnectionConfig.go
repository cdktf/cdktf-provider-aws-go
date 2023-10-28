// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkmanagerConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#connected_device_id NetworkmanagerConnection#connected_device_id}.
	ConnectedDeviceId *string `field:"required" json:"connectedDeviceId" yaml:"connectedDeviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#device_id NetworkmanagerConnection#device_id}.
	DeviceId *string `field:"required" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#global_network_id NetworkmanagerConnection#global_network_id}.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#connected_link_id NetworkmanagerConnection#connected_link_id}.
	ConnectedLinkId *string `field:"optional" json:"connectedLinkId" yaml:"connectedLinkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#description NetworkmanagerConnection#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#id NetworkmanagerConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#link_id NetworkmanagerConnection#link_id}.
	LinkId *string `field:"optional" json:"linkId" yaml:"linkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#tags NetworkmanagerConnection#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#tags_all NetworkmanagerConnection#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/networkmanager_connection#timeouts NetworkmanagerConnection#timeouts}
	Timeouts *NetworkmanagerConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

