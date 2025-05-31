// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lexv2modelsbotversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Lexv2ModelsBotVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/lexv2models_bot_version#bot_id Lexv2ModelsBotVersion#bot_id}.
	BotId *string `field:"required" json:"botId" yaml:"botId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/lexv2models_bot_version#locale_specification Lexv2ModelsBotVersion#locale_specification}.
	LocaleSpecification interface{} `field:"required" json:"localeSpecification" yaml:"localeSpecification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/lexv2models_bot_version#bot_version Lexv2ModelsBotVersion#bot_version}.
	BotVersion *string `field:"optional" json:"botVersion" yaml:"botVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/lexv2models_bot_version#description Lexv2ModelsBotVersion#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/lexv2models_bot_version#timeouts Lexv2ModelsBotVersion#timeouts}
	Timeouts *Lexv2ModelsBotVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

