// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesv2dedicatedippool


type Sesv2DedicatedIpPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/sesv2_dedicated_ip_pool#create Sesv2DedicatedIpPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/sesv2_dedicated_ip_pool#delete Sesv2DedicatedIpPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.21.0/docs/resources/sesv2_dedicated_ip_pool#update Sesv2DedicatedIpPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

