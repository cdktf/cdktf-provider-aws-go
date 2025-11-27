// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagercontrol

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerControlConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#name AuditmanagerControl#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#action_plan_instructions AuditmanagerControl#action_plan_instructions}.
	ActionPlanInstructions *string `field:"optional" json:"actionPlanInstructions" yaml:"actionPlanInstructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#action_plan_title AuditmanagerControl#action_plan_title}.
	ActionPlanTitle *string `field:"optional" json:"actionPlanTitle" yaml:"actionPlanTitle"`
	// control_mapping_sources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#control_mapping_sources AuditmanagerControl#control_mapping_sources}
	ControlMappingSources interface{} `field:"optional" json:"controlMappingSources" yaml:"controlMappingSources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#description AuditmanagerControl#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#region AuditmanagerControl#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#tags AuditmanagerControl#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/auditmanager_control#testing_information AuditmanagerControl#testing_information}.
	TestingInformation *string `field:"optional" json:"testingInformation" yaml:"testingInformation"`
}

