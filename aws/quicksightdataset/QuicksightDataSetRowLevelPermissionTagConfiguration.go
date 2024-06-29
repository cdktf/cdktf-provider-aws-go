// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetRowLevelPermissionTagConfiguration struct {
	// tag_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/quicksight_data_set#tag_rules QuicksightDataSet#tag_rules}
	TagRules interface{} `field:"required" json:"tagRules" yaml:"tagRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.56.1/docs/resources/quicksight_data_set#status QuicksightDataSet#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

