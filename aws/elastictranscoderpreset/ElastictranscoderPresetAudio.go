// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpreset


type ElastictranscoderPresetAudio struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/elastictranscoder_preset#audio_packing_mode ElastictranscoderPreset#audio_packing_mode}.
	AudioPackingMode *string `field:"optional" json:"audioPackingMode" yaml:"audioPackingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/elastictranscoder_preset#bit_rate ElastictranscoderPreset#bit_rate}.
	BitRate *string `field:"optional" json:"bitRate" yaml:"bitRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/elastictranscoder_preset#channels ElastictranscoderPreset#channels}.
	Channels *string `field:"optional" json:"channels" yaml:"channels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/elastictranscoder_preset#codec ElastictranscoderPreset#codec}.
	Codec *string `field:"optional" json:"codec" yaml:"codec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.21.0/docs/resources/elastictranscoder_preset#sample_rate ElastictranscoderPreset#sample_rate}.
	SampleRate *string `field:"optional" json:"sampleRate" yaml:"sampleRate"`
}

