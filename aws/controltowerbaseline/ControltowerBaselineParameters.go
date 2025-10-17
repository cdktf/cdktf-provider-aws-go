// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package controltowerbaseline


type ControltowerBaselineParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/controltower_baseline#key ControltowerBaseline#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.17.0/docs/resources/controltower_baseline#value ControltowerBaseline#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

