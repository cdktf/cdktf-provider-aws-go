// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package timestreamwritetable


type TimestreamwriteTableSchemaCompositePartitionKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/timestreamwrite_table#type TimestreamwriteTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/timestreamwrite_table#enforcement_in_record TimestreamwriteTable#enforcement_in_record}.
	EnforcementInRecord *string `field:"optional" json:"enforcementInRecord" yaml:"enforcementInRecord"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.13.0/docs/resources/timestreamwrite_table#name TimestreamwriteTable#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

