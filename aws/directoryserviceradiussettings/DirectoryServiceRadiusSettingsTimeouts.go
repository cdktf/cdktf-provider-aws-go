// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package directoryserviceradiussettings


type DirectoryServiceRadiusSettingsTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/directory_service_radius_settings#create DirectoryServiceRadiusSettings#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.69.0/docs/resources/directory_service_radius_settings#update DirectoryServiceRadiusSettings#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

