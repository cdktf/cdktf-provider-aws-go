// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesreceiptrule


type SesReceiptRuleAddHeaderAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ses_receipt_rule#header_name SesReceiptRule#header_name}.
	HeaderName *string `field:"required" json:"headerName" yaml:"headerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ses_receipt_rule#header_value SesReceiptRule#header_value}.
	HeaderValue *string `field:"required" json:"headerValue" yaml:"headerValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/ses_receipt_rule#position SesReceiptRule#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
}

