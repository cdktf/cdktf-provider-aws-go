// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package amicopy


type AmiCopyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/ami_copy#create AmiCopy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/ami_copy#delete AmiCopy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.1/docs/resources/ami_copy#update AmiCopy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

