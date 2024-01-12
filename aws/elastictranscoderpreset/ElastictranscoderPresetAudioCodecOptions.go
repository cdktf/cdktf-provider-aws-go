// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpreset


type ElastictranscoderPresetAudioCodecOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/elastictranscoder_preset#bit_depth ElastictranscoderPreset#bit_depth}.
	BitDepth *string `field:"optional" json:"bitDepth" yaml:"bitDepth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/elastictranscoder_preset#bit_order ElastictranscoderPreset#bit_order}.
	BitOrder *string `field:"optional" json:"bitOrder" yaml:"bitOrder"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/elastictranscoder_preset#profile ElastictranscoderPreset#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.32.0/docs/resources/elastictranscoder_preset#signed ElastictranscoderPreset#signed}.
	Signed *string `field:"optional" json:"signed" yaml:"signed"`
}

