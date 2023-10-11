// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lb


type LbTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/lb#create Lb#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/lb#delete Lb#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.20.1/docs/resources/lb#update Lb#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

