// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssmquicksetupconfigurationmanager

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SsmquicksetupConfigurationManagerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ssmquicksetup_configuration_manager#name SsmquicksetupConfigurationManager#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// configuration_definition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ssmquicksetup_configuration_manager#configuration_definition SsmquicksetupConfigurationManager#configuration_definition}
	ConfigurationDefinition interface{} `field:"optional" json:"configurationDefinition" yaml:"configurationDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ssmquicksetup_configuration_manager#description SsmquicksetupConfigurationManager#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ssmquicksetup_configuration_manager#tags SsmquicksetupConfigurationManager#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/ssmquicksetup_configuration_manager#timeouts SsmquicksetupConfigurationManager#timeouts}
	Timeouts *SsmquicksetupConfigurationManagerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

