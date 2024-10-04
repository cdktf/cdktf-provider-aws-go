// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcipampoolcidr


type VpcIpamPoolCidrCidrAuthorizationContext struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/vpc_ipam_pool_cidr#message VpcIpamPoolCidr#message}.
	Message *string `field:"optional" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.70.0/docs/resources/vpc_ipam_pool_cidr#signature VpcIpamPoolCidr#signature}.
	Signature *string `field:"optional" json:"signature" yaml:"signature"`
}

