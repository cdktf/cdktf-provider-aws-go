// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmschannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PinpointSmsChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#application_id PinpointSmsChannel#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#enabled PinpointSmsChannel#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#id PinpointSmsChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#region PinpointSmsChannel#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#sender_id PinpointSmsChannel#sender_id}.
	SenderId *string `field:"optional" json:"senderId" yaml:"senderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/pinpoint_sms_channel#short_code PinpointSmsChannel#short_code}.
	ShortCode *string `field:"optional" json:"shortCode" yaml:"shortCode"`
}

