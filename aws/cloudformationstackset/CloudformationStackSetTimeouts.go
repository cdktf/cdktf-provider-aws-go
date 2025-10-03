// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset


type CloudformationStackSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.15.0/docs/resources/cloudformation_stack_set#update CloudformationStackSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

