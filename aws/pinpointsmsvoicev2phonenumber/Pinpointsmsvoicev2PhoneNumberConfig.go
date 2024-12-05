// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2phonenumber

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Pinpointsmsvoicev2PhoneNumberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#iso_country_code Pinpointsmsvoicev2PhoneNumber#iso_country_code}.
	IsoCountryCode *string `field:"required" json:"isoCountryCode" yaml:"isoCountryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#message_type Pinpointsmsvoicev2PhoneNumber#message_type}.
	MessageType *string `field:"required" json:"messageType" yaml:"messageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#number_capabilities Pinpointsmsvoicev2PhoneNumber#number_capabilities}.
	NumberCapabilities *[]*string `field:"required" json:"numberCapabilities" yaml:"numberCapabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#number_type Pinpointsmsvoicev2PhoneNumber#number_type}.
	NumberType *string `field:"required" json:"numberType" yaml:"numberType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#deletion_protection_enabled Pinpointsmsvoicev2PhoneNumber#deletion_protection_enabled}.
	DeletionProtectionEnabled interface{} `field:"optional" json:"deletionProtectionEnabled" yaml:"deletionProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#opt_out_list_name Pinpointsmsvoicev2PhoneNumber#opt_out_list_name}.
	OptOutListName *string `field:"optional" json:"optOutListName" yaml:"optOutListName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#registration_id Pinpointsmsvoicev2PhoneNumber#registration_id}.
	RegistrationId *string `field:"optional" json:"registrationId" yaml:"registrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#self_managed_opt_outs_enabled Pinpointsmsvoicev2PhoneNumber#self_managed_opt_outs_enabled}.
	SelfManagedOptOutsEnabled interface{} `field:"optional" json:"selfManagedOptOutsEnabled" yaml:"selfManagedOptOutsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#tags Pinpointsmsvoicev2PhoneNumber#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#timeouts Pinpointsmsvoicev2PhoneNumber#timeouts}
	Timeouts *Pinpointsmsvoicev2PhoneNumberTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#two_way_channel_arn Pinpointsmsvoicev2PhoneNumber#two_way_channel_arn}.
	TwoWayChannelArn *string `field:"optional" json:"twoWayChannelArn" yaml:"twoWayChannelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/pinpointsmsvoicev2_phone_number#two_way_channel_enabled Pinpointsmsvoicev2PhoneNumber#two_way_channel_enabled}.
	TwoWayChannelEnabled interface{} `field:"optional" json:"twoWayChannelEnabled" yaml:"twoWayChannelEnabled"`
}

