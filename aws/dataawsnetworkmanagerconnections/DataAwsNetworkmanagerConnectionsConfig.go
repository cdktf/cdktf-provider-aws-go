// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsnetworkmanagerconnections

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsNetworkmanagerConnectionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/networkmanager_connections#global_network_id DataAwsNetworkmanagerConnections#global_network_id}.
	GlobalNetworkId *string `field:"required" json:"globalNetworkId" yaml:"globalNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/networkmanager_connections#device_id DataAwsNetworkmanagerConnections#device_id}.
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/networkmanager_connections#id DataAwsNetworkmanagerConnections#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/data-sources/networkmanager_connections#tags DataAwsNetworkmanagerConnections#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

