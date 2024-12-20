// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel


type MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettings struct {
	// hls_akamai_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/medialive_channel#hls_akamai_settings MedialiveChannel#hls_akamai_settings}
	HlsAkamaiSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsAkamaiSettings `field:"optional" json:"hlsAkamaiSettings" yaml:"hlsAkamaiSettings"`
	// hls_basic_put_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/medialive_channel#hls_basic_put_settings MedialiveChannel#hls_basic_put_settings}
	HlsBasicPutSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsBasicPutSettings `field:"optional" json:"hlsBasicPutSettings" yaml:"hlsBasicPutSettings"`
	// hls_media_store_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/medialive_channel#hls_media_store_settings MedialiveChannel#hls_media_store_settings}
	HlsMediaStoreSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsMediaStoreSettings `field:"optional" json:"hlsMediaStoreSettings" yaml:"hlsMediaStoreSettings"`
	// hls_s3_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/medialive_channel#hls_s3_settings MedialiveChannel#hls_s3_settings}
	HlsS3Settings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsS3Settings `field:"optional" json:"hlsS3Settings" yaml:"hlsS3Settings"`
	// hls_webdav_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.82.1/docs/resources/medialive_channel#hls_webdav_settings MedialiveChannel#hls_webdav_settings}
	HlsWebdavSettings *MedialiveChannelEncoderSettingsOutputGroupsOutputGroupSettingsHlsGroupSettingsHlsCdnSettingsHlsWebdavSettings `field:"optional" json:"hlsWebdavSettings" yaml:"hlsWebdavSettings"`
}

