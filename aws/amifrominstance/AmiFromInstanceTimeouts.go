// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amifrominstance


type AmiFromInstanceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ami_from_instance#create AmiFromInstance#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ami_from_instance#delete AmiFromInstance#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/ami_from_instance#update AmiFromInstance#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

