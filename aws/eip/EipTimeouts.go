// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eip


type EipTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/eip#delete Eip#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/eip#read Eip#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.72.1/docs/resources/eip#update Eip#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

