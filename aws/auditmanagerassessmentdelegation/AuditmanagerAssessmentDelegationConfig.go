// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auditmanagerassessmentdelegation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AuditmanagerAssessmentDelegationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#assessment_id AuditmanagerAssessmentDelegation#assessment_id}.
	AssessmentId *string `field:"required" json:"assessmentId" yaml:"assessmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#control_set_id AuditmanagerAssessmentDelegation#control_set_id}.
	ControlSetId *string `field:"required" json:"controlSetId" yaml:"controlSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#role_arn AuditmanagerAssessmentDelegation#role_arn}.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#role_type AuditmanagerAssessmentDelegation#role_type}.
	RoleType *string `field:"required" json:"roleType" yaml:"roleType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#comment AuditmanagerAssessmentDelegation#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/auditmanager_assessment_delegation#region AuditmanagerAssessmentDelegation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

