// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package snssmspreferences

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SnsSmsPreferencesConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#default_sender_id SnsSmsPreferences#default_sender_id}.
	DefaultSenderId *string `field:"optional" json:"defaultSenderId" yaml:"defaultSenderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#default_sms_type SnsSmsPreferences#default_sms_type}.
	DefaultSmsType *string `field:"optional" json:"defaultSmsType" yaml:"defaultSmsType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#delivery_status_iam_role_arn SnsSmsPreferences#delivery_status_iam_role_arn}.
	DeliveryStatusIamRoleArn *string `field:"optional" json:"deliveryStatusIamRoleArn" yaml:"deliveryStatusIamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#delivery_status_success_sampling_rate SnsSmsPreferences#delivery_status_success_sampling_rate}.
	DeliveryStatusSuccessSamplingRate *string `field:"optional" json:"deliveryStatusSuccessSamplingRate" yaml:"deliveryStatusSuccessSamplingRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#id SnsSmsPreferences#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#monthly_spend_limit SnsSmsPreferences#monthly_spend_limit}.
	MonthlySpendLimit *float64 `field:"optional" json:"monthlySpendLimit" yaml:"monthlySpendLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.59.0/docs/resources/sns_sms_preferences#usage_report_s3_bucket SnsSmsPreferences#usage_report_s3_bucket}.
	UsageReportS3Bucket *string `field:"optional" json:"usageReportS3Bucket" yaml:"usageReportS3Bucket"`
}

