// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsAkamaiSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#connection_retry_interval MedialiveChannel#connection_retry_interval}.
	ConnectionRetryInterval *float64 `field:"optional" json:"connectionRetryInterval" yaml:"connectionRetryInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#filecache_duration MedialiveChannel#filecache_duration}.
	FilecacheDuration *float64 `field:"optional" json:"filecacheDuration" yaml:"filecacheDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#http_transfer_mode MedialiveChannel#http_transfer_mode}.
	HttpTransferMode *string `field:"optional" json:"httpTransferMode" yaml:"httpTransferMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#num_retries MedialiveChannel#num_retries}.
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#restart_delay MedialiveChannel#restart_delay}.
	RestartDelay *float64 `field:"optional" json:"restartDelay" yaml:"restartDelay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#salt MedialiveChannel#salt}.
	Salt *string `field:"optional" json:"salt" yaml:"salt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/medialive_channel#token MedialiveChannel#token}.
	Token *string `field:"optional" json:"token" yaml:"token"`
}

