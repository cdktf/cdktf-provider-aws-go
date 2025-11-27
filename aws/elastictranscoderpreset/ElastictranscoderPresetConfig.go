// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package elastictranscoderpreset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElastictranscoderPresetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#container ElastictranscoderPreset#container}.
	Container *string `field:"required" json:"container" yaml:"container"`
	// audio block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#audio ElastictranscoderPreset#audio}
	Audio *ElastictranscoderPresetAudio `field:"optional" json:"audio" yaml:"audio"`
	// audio_codec_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#audio_codec_options ElastictranscoderPreset#audio_codec_options}
	AudioCodecOptions *ElastictranscoderPresetAudioCodecOptions `field:"optional" json:"audioCodecOptions" yaml:"audioCodecOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#description ElastictranscoderPreset#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#id ElastictranscoderPreset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#name ElastictranscoderPreset#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#region ElastictranscoderPreset#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// thumbnails block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#thumbnails ElastictranscoderPreset#thumbnails}
	Thumbnails *ElastictranscoderPresetThumbnails `field:"optional" json:"thumbnails" yaml:"thumbnails"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#type ElastictranscoderPreset#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// video block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#video ElastictranscoderPreset#video}
	Video *ElastictranscoderPresetVideo `field:"optional" json:"video" yaml:"video"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#video_codec_options ElastictranscoderPreset#video_codec_options}.
	VideoCodecOptions *map[string]*string `field:"optional" json:"videoCodecOptions" yaml:"videoCodecOptions"`
	// video_watermarks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/elastictranscoder_preset#video_watermarks ElastictranscoderPreset#video_watermarks}
	VideoWatermarks interface{} `field:"optional" json:"videoWatermarks" yaml:"videoWatermarks"`
}

