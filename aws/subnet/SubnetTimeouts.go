// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subnet


type SubnetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/subnet#create Subnet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/subnet#delete Subnet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

