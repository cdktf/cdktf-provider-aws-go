// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataawsauditmanagerframework

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAwsAuditmanagerFrameworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/data-sources/auditmanager_framework#framework_type DataAwsAuditmanagerFramework#framework_type}.
	FrameworkType *string `field:"required" json:"frameworkType" yaml:"frameworkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/data-sources/auditmanager_framework#name DataAwsAuditmanagerFramework#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// control_sets block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.51.0/docs/data-sources/auditmanager_framework#control_sets DataAwsAuditmanagerFramework#control_sets}
	ControlSets interface{} `field:"optional" json:"controlSets" yaml:"controlSets"`
}

