// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerframeworkshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerFrameworkShareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/auditmanager_framework_share#destination_account AuditmanagerFrameworkShare#destination_account}.
	DestinationAccount *string `field:"required" json:"destinationAccount" yaml:"destinationAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/auditmanager_framework_share#destination_region AuditmanagerFrameworkShare#destination_region}.
	DestinationRegion *string `field:"required" json:"destinationRegion" yaml:"destinationRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/auditmanager_framework_share#framework_id AuditmanagerFrameworkShare#framework_id}.
	FrameworkId *string `field:"required" json:"frameworkId" yaml:"frameworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.26.0/docs/resources/auditmanager_framework_share#comment AuditmanagerFrameworkShare#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

