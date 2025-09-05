// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimevoiceconnectorterminationcredentials


type ChimeVoiceConnectorTerminationCredentialsCredentials struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chime_voice_connector_termination_credentials#password ChimeVoiceConnectorTerminationCredentials#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/chime_voice_connector_termination_credentials#username ChimeVoiceConnectorTerminationCredentials#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}

