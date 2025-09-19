// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimevoiceconnectororigination


type ChimeVoiceConnectorOriginationRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/chime_voice_connector_origination#host ChimeVoiceConnectorOrigination#host}.
	Host *string `field:"required" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/chime_voice_connector_origination#priority ChimeVoiceConnectorOrigination#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/chime_voice_connector_origination#protocol ChimeVoiceConnectorOrigination#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/chime_voice_connector_origination#weight ChimeVoiceConnectorOrigination#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.14.0/docs/resources/chime_voice_connector_origination#port ChimeVoiceConnectorOrigination#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

