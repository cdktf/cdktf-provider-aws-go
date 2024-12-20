// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtrail


type CloudtrailAdvancedEventSelector struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/cloudtrail#field_selector Cloudtrail#field_selector}
	FieldSelector interface{} `field:"required" json:"fieldSelector" yaml:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/cloudtrail#name Cloudtrail#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

