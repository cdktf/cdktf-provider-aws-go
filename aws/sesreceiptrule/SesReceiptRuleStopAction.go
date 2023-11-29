// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesreceiptrule


type SesReceiptRuleStopAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ses_receipt_rule#position SesReceiptRule#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ses_receipt_rule#scope SesReceiptRule#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.28.0/docs/resources/ses_receipt_rule#topic_arn SesReceiptRule#topic_arn}.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

