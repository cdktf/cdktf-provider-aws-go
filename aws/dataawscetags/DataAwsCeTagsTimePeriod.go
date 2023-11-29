// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawscetags


type DataAwsCeTagsTimePeriod struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/data-sources/ce_tags#end DataAwsCeTags#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/data-sources/ce_tags#start DataAwsCeTags#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

