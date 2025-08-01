// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ami


type AmiTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ami#create Ami#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ami#delete Ami#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/ami#update Ami#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

