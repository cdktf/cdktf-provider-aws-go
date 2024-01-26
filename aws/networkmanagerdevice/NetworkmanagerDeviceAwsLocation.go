// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerdevice


type NetworkmanagerDeviceAwsLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/networkmanager_device#subnet_arn NetworkmanagerDevice#subnet_arn}.
	SubnetArn *string `field:"optional" json:"subnetArn" yaml:"subnetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.34.0/docs/resources/networkmanager_device#zone NetworkmanagerDevice#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

