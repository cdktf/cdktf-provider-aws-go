// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpoint


type VpcEndpointSubnetConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/vpc_endpoint#ipv4 VpcEndpoint#ipv4}.
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/vpc_endpoint#ipv6 VpcEndpoint#ipv6}.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/vpc_endpoint#subnet_id VpcEndpoint#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

