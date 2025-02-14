// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerAssessmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#framework_id AuditmanagerAssessment#framework_id}.
	FrameworkId *string `field:"required" json:"frameworkId" yaml:"frameworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#name AuditmanagerAssessment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#roles AuditmanagerAssessment#roles}.
	Roles interface{} `field:"required" json:"roles" yaml:"roles"`
	// assessment_reports_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#assessment_reports_destination AuditmanagerAssessment#assessment_reports_destination}
	AssessmentReportsDestination interface{} `field:"optional" json:"assessmentReportsDestination" yaml:"assessmentReportsDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#description AuditmanagerAssessment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#scope AuditmanagerAssessment#scope}
	Scope interface{} `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.87.0/docs/resources/auditmanager_assessment#tags AuditmanagerAssessment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

