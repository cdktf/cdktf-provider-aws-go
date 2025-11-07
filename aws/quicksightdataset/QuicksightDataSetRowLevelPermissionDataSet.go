// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightdataset


type QuicksightDataSetRowLevelPermissionDataSet struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_data_set#arn QuicksightDataSet#arn}.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_data_set#permission_policy QuicksightDataSet#permission_policy}.
	PermissionPolicy *string `field:"required" json:"permissionPolicy" yaml:"permissionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_data_set#format_version QuicksightDataSet#format_version}.
	FormatVersion *string `field:"optional" json:"formatVersion" yaml:"formatVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_data_set#namespace QuicksightDataSet#namespace}.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.20.0/docs/resources/quicksight_data_set#status QuicksightDataSet#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

