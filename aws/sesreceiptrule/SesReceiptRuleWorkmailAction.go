// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sesreceiptrule


type SesReceiptRuleWorkmailAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/ses_receipt_rule#organization_arn SesReceiptRule#organization_arn}.
	OrganizationArn *string `field:"required" json:"organizationArn" yaml:"organizationArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/ses_receipt_rule#position SesReceiptRule#position}.
	Position *float64 `field:"required" json:"position" yaml:"position"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.3.0/docs/resources/ses_receipt_rule#topic_arn SesReceiptRule#topic_arn}.
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

