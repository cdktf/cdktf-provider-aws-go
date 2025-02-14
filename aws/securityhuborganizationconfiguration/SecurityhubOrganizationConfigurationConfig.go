// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhuborganizationconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubOrganizationConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/securityhub_organization_configuration#auto_enable SecurityhubOrganizationConfiguration#auto_enable}.
	AutoEnable interface{} `field:"required" json:"autoEnable" yaml:"autoEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/securityhub_organization_configuration#auto_enable_standards SecurityhubOrganizationConfiguration#auto_enable_standards}.
	AutoEnableStandards *string `field:"optional" json:"autoEnableStandards" yaml:"autoEnableStandards"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/securityhub_organization_configuration#id SecurityhubOrganizationConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// organization_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/securityhub_organization_configuration#organization_configuration SecurityhubOrganizationConfiguration#organization_configuration}
	OrganizationConfiguration *SecurityhubOrganizationConfigurationOrganizationConfiguration `field:"optional" json:"organizationConfiguration" yaml:"organizationConfiguration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/securityhub_organization_configuration#timeouts SecurityhubOrganizationConfiguration#timeouts}
	Timeouts *SecurityhubOrganizationConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

