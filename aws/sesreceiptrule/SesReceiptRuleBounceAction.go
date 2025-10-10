// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesreceiptrule


type SesReceiptRuleBounceAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#message SesReceiptRule#message}.
	Message *string `field:"required" json:"message" yaml:"message"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#position SesReceiptRule#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#sender SesReceiptRule#sender}.
	Sender *string `field:"required" json:"sender" yaml:"sender"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#smtp_reply_code SesReceiptRule#smtp_reply_code}.
	SmtpReplyCode *string `field:"required" json:"smtpReplyCode" yaml:"smtpReplyCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#status_code SesReceiptRule#status_code}.
	StatusCode *string `field:"optional" json:"statusCode" yaml:"statusCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.16.0/docs/resources/ses_receipt_rule#topic_arn SesReceiptRule#topic_arn}.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

