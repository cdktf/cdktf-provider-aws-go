// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaceswebsessionloggerassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspaceswebSessionLoggerAssociationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/workspacesweb_session_logger_association#portal_arn WorkspaceswebSessionLoggerAssociation#portal_arn}.
	PortalArn *string `field:"required" json:"portalArn" yaml:"portalArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/workspacesweb_session_logger_association#session_logger_arn WorkspaceswebSessionLoggerAssociation#session_logger_arn}.
	SessionLoggerArn *string `field:"required" json:"sessionLoggerArn" yaml:"sessionLoggerArn"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.12.0/docs/resources/workspacesweb_session_logger_association#region WorkspaceswebSessionLoggerAssociation#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
}

