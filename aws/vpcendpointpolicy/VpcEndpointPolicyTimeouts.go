// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcendpointpolicy


type VpcEndpointPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/vpc_endpoint_policy#create VpcEndpointPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.0/docs/resources/vpc_endpoint_policy#delete VpcEndpointPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

