// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelDestinationsSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#password_param MedialiveChannel#password_param}.
	PasswordParam *string `field:"optional" json:"passwordParam" yaml:"passwordParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#stream_name MedialiveChannel#stream_name}.
	StreamName *string `field:"optional" json:"streamName" yaml:"streamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#url MedialiveChannel#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/medialive_channel#username MedialiveChannel#username}.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

