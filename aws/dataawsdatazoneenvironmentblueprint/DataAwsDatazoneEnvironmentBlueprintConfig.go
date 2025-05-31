// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsdatazoneenvironmentblueprint

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsDatazoneEnvironmentBlueprintConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/datazone_environment_blueprint#domain_id DataAwsDatazoneEnvironmentBlueprint#domain_id}.
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/datazone_environment_blueprint#managed DataAwsDatazoneEnvironmentBlueprint#managed}.
	Managed interface{} `field:"required" json:"managed" yaml:"managed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.99.1/docs/data-sources/datazone_environment_blueprint#name DataAwsDatazoneEnvironmentBlueprint#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

