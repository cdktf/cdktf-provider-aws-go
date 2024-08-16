// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointapnsvoipchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PinpointApnsVoipChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#application_id PinpointApnsVoipChannel#application_id}.
	ApplicationId *string `field:"required" json:"applicationId" yaml:"applicationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#bundle_id PinpointApnsVoipChannel#bundle_id}.
	BundleId *string `field:"optional" json:"bundleId" yaml:"bundleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#certificate PinpointApnsVoipChannel#certificate}.
	Certificate *string `field:"optional" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#default_authentication_method PinpointApnsVoipChannel#default_authentication_method}.
	DefaultAuthenticationMethod *string `field:"optional" json:"defaultAuthenticationMethod" yaml:"defaultAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#enabled PinpointApnsVoipChannel#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#id PinpointApnsVoipChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#private_key PinpointApnsVoipChannel#private_key}.
	PrivateKey *string `field:"optional" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#team_id PinpointApnsVoipChannel#team_id}.
	TeamId *string `field:"optional" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#token_key PinpointApnsVoipChannel#token_key}.
	TokenKey *string `field:"optional" json:"tokenKey" yaml:"tokenKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.63.0/docs/resources/pinpoint_apns_voip_channel#token_key_id PinpointApnsVoipChannel#token_key_id}.
	TokenKeyId *string `field:"optional" json:"tokenKeyId" yaml:"tokenKeyId"`
}

