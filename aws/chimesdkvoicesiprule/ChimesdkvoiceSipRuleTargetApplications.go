// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimesdkvoicesiprule


type ChimesdkvoiceSipRuleTargetApplications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkvoice_sip_rule#aws_region ChimesdkvoiceSipRule#aws_region}.
	AwsRegion *string `field:"required" json:"awsRegion" yaml:"awsRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkvoice_sip_rule#priority ChimesdkvoiceSipRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.45.0/docs/resources/chimesdkvoice_sip_rule#sip_media_application_id ChimesdkvoiceSipRule#sip_media_application_id}.
	SipMediaApplicationId *string `field:"required" json:"sipMediaApplicationId" yaml:"sipMediaApplicationId"`
}

