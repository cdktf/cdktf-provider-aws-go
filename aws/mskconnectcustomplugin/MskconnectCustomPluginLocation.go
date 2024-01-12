// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mskconnectcustomplugin


type MskconnectCustomPluginLocation struct {
	// s3 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/mskconnect_custom_plugin#s3 MskconnectCustomPlugin#s3}
	S3 *MskconnectCustomPluginLocationS3 `field:"required" json:"s3" yaml:"s3"`
}

