// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpreset


type ElastictranscoderPresetVideoWatermarks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#horizontal_align ElastictranscoderPreset#horizontal_align}.
	HorizontalAlign *string `field:"optional" json:"horizontalAlign" yaml:"horizontalAlign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#horizontal_offset ElastictranscoderPreset#horizontal_offset}.
	HorizontalOffset *string `field:"optional" json:"horizontalOffset" yaml:"horizontalOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#id ElastictranscoderPreset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#max_height ElastictranscoderPreset#max_height}.
	MaxHeight *string `field:"optional" json:"maxHeight" yaml:"maxHeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#max_width ElastictranscoderPreset#max_width}.
	MaxWidth *string `field:"optional" json:"maxWidth" yaml:"maxWidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#opacity ElastictranscoderPreset#opacity}.
	Opacity *string `field:"optional" json:"opacity" yaml:"opacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#sizing_policy ElastictranscoderPreset#sizing_policy}.
	SizingPolicy *string `field:"optional" json:"sizingPolicy" yaml:"sizingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#target ElastictranscoderPreset#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#vertical_align ElastictranscoderPreset#vertical_align}.
	VerticalAlign *string `field:"optional" json:"verticalAlign" yaml:"verticalAlign"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.78.0/docs/resources/elastictranscoder_preset#vertical_offset ElastictranscoderPreset#vertical_offset}.
	VerticalOffset *string `field:"optional" json:"verticalOffset" yaml:"verticalOffset"`
}

