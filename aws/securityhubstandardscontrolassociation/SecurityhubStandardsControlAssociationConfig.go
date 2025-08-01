// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securityhubstandardscontrolassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityhubStandardsControlAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_standards_control_association#association_status SecurityhubStandardsControlAssociation#association_status}.
	AssociationStatus *string `field:"required" json:"associationStatus" yaml:"associationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_standards_control_association#security_control_id SecurityhubStandardsControlAssociation#security_control_id}.
	SecurityControlId *string `field:"required" json:"securityControlId" yaml:"securityControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_standards_control_association#standards_arn SecurityhubStandardsControlAssociation#standards_arn}.
	StandardsArn *string `field:"required" json:"standardsArn" yaml:"standardsArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_standards_control_association#region SecurityhubStandardsControlAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.7.0/docs/resources/securityhub_standards_control_association#updated_reason SecurityhubStandardsControlAssociation#updated_reason}.
	UpdatedReason *string `field:"optional" json:"updatedReason" yaml:"updatedReason"`
}

