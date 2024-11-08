// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package worklinkfleet


type WorklinkFleetIdentityProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#saml_metadata WorklinkFleet#saml_metadata}.
	SamlMetadata *string `field:"required" json:"samlMetadata" yaml:"samlMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.0/docs/resources/worklink_fleet#type WorklinkFleet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

