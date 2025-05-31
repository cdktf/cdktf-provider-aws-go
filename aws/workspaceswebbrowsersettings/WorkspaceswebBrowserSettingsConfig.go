// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebbrowsersettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceswebBrowserSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/workspacesweb_browser_settings#browser_policy WorkspaceswebBrowserSettings#browser_policy}.
	BrowserPolicy *string `field:"required" json:"browserPolicy" yaml:"browserPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/workspacesweb_browser_settings#additional_encryption_context WorkspaceswebBrowserSettings#additional_encryption_context}.
	AdditionalEncryptionContext *map[string]*string `field:"optional" json:"additionalEncryptionContext" yaml:"additionalEncryptionContext"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/workspacesweb_browser_settings#customer_managed_key WorkspaceswebBrowserSettings#customer_managed_key}.
	CustomerManagedKey *string `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/resources/workspacesweb_browser_settings#tags WorkspaceswebBrowserSettings#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

