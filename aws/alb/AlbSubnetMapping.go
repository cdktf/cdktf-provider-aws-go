// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package alb


type AlbSubnetMapping struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb#subnet_id Alb#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb#allocation_id Alb#allocation_id}.
	AllocationId *string `field:"optional" json:"allocationId" yaml:"allocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb#ipv6_address Alb#ipv6_address}.
	Ipv6Address *string `field:"optional" json:"ipv6Address" yaml:"ipv6Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/alb#private_ipv4_address Alb#private_ipv4_address}.
	PrivateIpv4Address *string `field:"optional" json:"privateIpv4Address" yaml:"privateIpv4Address"`
}

