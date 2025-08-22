// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package quicksightaccountsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type QuicksightAccountSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_account_settings#aws_account_id QuicksightAccountSettings#aws_account_id}.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_account_settings#default_namespace QuicksightAccountSettings#default_namespace}.
	DefaultNamespace *string `field:"optional" json:"defaultNamespace" yaml:"defaultNamespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_account_settings#termination_protection_enabled QuicksightAccountSettings#termination_protection_enabled}.
	TerminationProtectionEnabled interface{} `field:"optional" json:"terminationProtectionEnabled" yaml:"terminationProtectionEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.10.0/docs/resources/quicksight_account_settings#timeouts QuicksightAccountSettings#timeouts}
	Timeouts *QuicksightAccountSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

