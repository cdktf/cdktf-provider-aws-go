// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chimevoiceconnectorstreaming

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChimeVoiceConnectorStreamingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#data_retention ChimeVoiceConnectorStreaming#data_retention}.
	DataRetention *float64 `field:"required" json:"dataRetention" yaml:"dataRetention"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#voice_connector_id ChimeVoiceConnectorStreaming#voice_connector_id}.
	VoiceConnectorId *string `field:"required" json:"voiceConnectorId" yaml:"voiceConnectorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#disabled ChimeVoiceConnectorStreaming#disabled}.
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#id ChimeVoiceConnectorStreaming#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// media_insights_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#media_insights_configuration ChimeVoiceConnectorStreaming#media_insights_configuration}
	MediaInsightsConfiguration *ChimeVoiceConnectorStreamingMediaInsightsConfiguration `field:"optional" json:"mediaInsightsConfiguration" yaml:"mediaInsightsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.81.0/docs/resources/chime_voice_connector_streaming#streaming_notification_targets ChimeVoiceConnectorStreaming#streaming_notification_targets}.
	StreamingNotificationTargets *[]*string `field:"optional" json:"streamingNotificationTargets" yaml:"streamingNotificationTargets"`
}

