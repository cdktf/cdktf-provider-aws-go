// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudtrail


type CloudtrailEventSelectorDataResource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cloudtrail#type Cloudtrail#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/cloudtrail#values Cloudtrail#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

