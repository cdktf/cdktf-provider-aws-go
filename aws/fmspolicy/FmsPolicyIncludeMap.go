// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fmspolicy


type FmsPolicyIncludeMap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fms_policy#account FmsPolicy#account}.
	Account *[]*string `field:"optional" json:"account" yaml:"account"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/fms_policy#orgunit FmsPolicy#orgunit}.
	Orgunit *[]*string `field:"optional" json:"orgunit" yaml:"orgunit"`
}

