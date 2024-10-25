// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connectinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ConnectInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#identity_management_type ConnectInstance#identity_management_type}.
	IdentityManagementType *string `field:"required" json:"identityManagementType" yaml:"identityManagementType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#inbound_calls_enabled ConnectInstance#inbound_calls_enabled}.
	InboundCallsEnabled interface{} `field:"required" json:"inboundCallsEnabled" yaml:"inboundCallsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#outbound_calls_enabled ConnectInstance#outbound_calls_enabled}.
	OutboundCallsEnabled interface{} `field:"required" json:"outboundCallsEnabled" yaml:"outboundCallsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#auto_resolve_best_voices_enabled ConnectInstance#auto_resolve_best_voices_enabled}.
	AutoResolveBestVoicesEnabled interface{} `field:"optional" json:"autoResolveBestVoicesEnabled" yaml:"autoResolveBestVoicesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#contact_flow_logs_enabled ConnectInstance#contact_flow_logs_enabled}.
	ContactFlowLogsEnabled interface{} `field:"optional" json:"contactFlowLogsEnabled" yaml:"contactFlowLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#contact_lens_enabled ConnectInstance#contact_lens_enabled}.
	ContactLensEnabled interface{} `field:"optional" json:"contactLensEnabled" yaml:"contactLensEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#directory_id ConnectInstance#directory_id}.
	DirectoryId *string `field:"optional" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#early_media_enabled ConnectInstance#early_media_enabled}.
	EarlyMediaEnabled interface{} `field:"optional" json:"earlyMediaEnabled" yaml:"earlyMediaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#id ConnectInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#instance_alias ConnectInstance#instance_alias}.
	InstanceAlias *string `field:"optional" json:"instanceAlias" yaml:"instanceAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#multi_party_conference_enabled ConnectInstance#multi_party_conference_enabled}.
	MultiPartyConferenceEnabled interface{} `field:"optional" json:"multiPartyConferenceEnabled" yaml:"multiPartyConferenceEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#tags ConnectInstance#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#tags_all ConnectInstance#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.73.0/docs/resources/connect_instance#timeouts ConnectInstance#timeouts}
	Timeouts *ConnectInstanceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

