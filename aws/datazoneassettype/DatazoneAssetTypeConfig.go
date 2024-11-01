// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datazoneassettype

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatazoneAssetTypeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#domain_identifier DatazoneAssetType#domain_identifier}.
	DomainIdentifier *string `field:"required" json:"domainIdentifier" yaml:"domainIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#name DatazoneAssetType#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#owning_project_identifier DatazoneAssetType#owning_project_identifier}.
	OwningProjectIdentifier *string `field:"required" json:"owningProjectIdentifier" yaml:"owningProjectIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#description DatazoneAssetType#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// forms_input block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#forms_input DatazoneAssetType#forms_input}
	FormsInput interface{} `field:"optional" json:"formsInput" yaml:"formsInput"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.74.0/docs/resources/datazone_asset_type#timeouts DatazoneAssetType#timeouts}
	Timeouts *DatazoneAssetTypeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

