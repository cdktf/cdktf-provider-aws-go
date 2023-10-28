// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstackset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationStackSetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#name CloudformationStackSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#administration_role_arn CloudformationStackSet#administration_role_arn}.
	AdministrationRoleArn *string `field:"optional" json:"administrationRoleArn" yaml:"administrationRoleArn"`
	// auto_deployment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#auto_deployment CloudformationStackSet#auto_deployment}
	AutoDeployment *CloudformationStackSetAutoDeployment `field:"optional" json:"autoDeployment" yaml:"autoDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#call_as CloudformationStackSet#call_as}.
	CallAs *string `field:"optional" json:"callAs" yaml:"callAs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#capabilities CloudformationStackSet#capabilities}.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#description CloudformationStackSet#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#execution_role_name CloudformationStackSet#execution_role_name}.
	ExecutionRoleName *string `field:"optional" json:"executionRoleName" yaml:"executionRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#id CloudformationStackSet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// managed_execution block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#managed_execution CloudformationStackSet#managed_execution}
	ManagedExecution *CloudformationStackSetManagedExecution `field:"optional" json:"managedExecution" yaml:"managedExecution"`
	// operation_preferences block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#operation_preferences CloudformationStackSet#operation_preferences}
	OperationPreferences *CloudformationStackSetOperationPreferences `field:"optional" json:"operationPreferences" yaml:"operationPreferences"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#parameters CloudformationStackSet#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#permission_model CloudformationStackSet#permission_model}.
	PermissionModel *string `field:"optional" json:"permissionModel" yaml:"permissionModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#tags CloudformationStackSet#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#tags_all CloudformationStackSet#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#template_body CloudformationStackSet#template_body}.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#template_url CloudformationStackSet#template_url}.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.23.1/docs/resources/cloudformation_stack_set#timeouts CloudformationStackSet#timeouts}
	Timeouts *CloudformationStackSetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

