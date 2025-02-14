// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsbot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsBotConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#idle_session_ttl_in_seconds Lexv2ModelsBot#idle_session_ttl_in_seconds}.
	IdleSessionTtlInSeconds *float64 `field:"required" json:"idleSessionTtlInSeconds" yaml:"idleSessionTtlInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#name Lexv2ModelsBot#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#role_arn Lexv2ModelsBot#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// data_privacy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#data_privacy Lexv2ModelsBot#data_privacy}
	DataPrivacy interface{} `field:"optional" json:"dataPrivacy" yaml:"dataPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#description Lexv2ModelsBot#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// members block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#members Lexv2ModelsBot#members}
	Members interface{} `field:"optional" json:"members" yaml:"members"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#tags Lexv2ModelsBot#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#test_bot_alias_tags Lexv2ModelsBot#test_bot_alias_tags}.
	TestBotAliasTags *map[string]*string `field:"optional" json:"testBotAliasTags" yaml:"testBotAliasTags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#timeouts Lexv2ModelsBot#timeouts}
	Timeouts *Lexv2ModelsBotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/lexv2models_bot#type Lexv2ModelsBot#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

