// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package medialivechannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#channel_class MedialiveChannel#channel_class}.
	ChannelClass *string `field:"required" json:"channelClass" yaml:"channelClass"`
	// destinations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#destinations MedialiveChannel#destinations}
	Destinations interface{} `field:"required" json:"destinations" yaml:"destinations"`
	// encoder_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#encoder_settings MedialiveChannel#encoder_settings}
	EncoderSettings *MedialiveChannelEncoderSettings `field:"required" json:"encoderSettings" yaml:"encoderSettings"`
	// input_attachments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#input_attachments MedialiveChannel#input_attachments}
	InputAttachments interface{} `field:"required" json:"inputAttachments" yaml:"inputAttachments"`
	// input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#input_specification MedialiveChannel#input_specification}
	InputSpecification *MedialiveChannelInputSpecification `field:"required" json:"inputSpecification" yaml:"inputSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#name MedialiveChannel#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cdi_input_specification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#cdi_input_specification MedialiveChannel#cdi_input_specification}
	CdiInputSpecification *MedialiveChannelCdiInputSpecification `field:"optional" json:"cdiInputSpecification" yaml:"cdiInputSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#id MedialiveChannel#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#log_level MedialiveChannel#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// maintenance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#maintenance MedialiveChannel#maintenance}
	Maintenance *MedialiveChannelMaintenance `field:"optional" json:"maintenance" yaml:"maintenance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#role_arn MedialiveChannel#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#start_channel MedialiveChannel#start_channel}.
	StartChannel interface{} `field:"optional" json:"startChannel" yaml:"startChannel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#tags MedialiveChannel#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#tags_all MedialiveChannel#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#timeouts MedialiveChannel#timeouts}
	Timeouts *MedialiveChannelTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vpc block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.95.0/docs/resources/medialive_channel#vpc MedialiveChannel#vpc}
	Vpc *MedialiveChannelVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

