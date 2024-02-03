// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudformationstack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CloudformationStackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#name CloudformationStack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#capabilities CloudformationStack#capabilities}.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#disable_rollback CloudformationStack#disable_rollback}.
	DisableRollback interface{} `field:"optional" json:"disableRollback" yaml:"disableRollback"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#iam_role_arn CloudformationStack#iam_role_arn}.
	IamRoleArn *string `field:"optional" json:"iamRoleArn" yaml:"iamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#id CloudformationStack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#notification_arns CloudformationStack#notification_arns}.
	NotificationArns *[]*string `field:"optional" json:"notificationArns" yaml:"notificationArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#on_failure CloudformationStack#on_failure}.
	OnFailure *string `field:"optional" json:"onFailure" yaml:"onFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#parameters CloudformationStack#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#policy_body CloudformationStack#policy_body}.
	PolicyBody *string `field:"optional" json:"policyBody" yaml:"policyBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#policy_url CloudformationStack#policy_url}.
	PolicyUrl *string `field:"optional" json:"policyUrl" yaml:"policyUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#tags CloudformationStack#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#tags_all CloudformationStack#tags_all}.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#template_body CloudformationStack#template_body}.
	TemplateBody *string `field:"optional" json:"templateBody" yaml:"templateBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#template_url CloudformationStack#template_url}.
	TemplateUrl *string `field:"optional" json:"templateUrl" yaml:"templateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#timeout_in_minutes CloudformationStack#timeout_in_minutes}.
	TimeoutInMinutes *float64 `field:"optional" json:"timeoutInMinutes" yaml:"timeoutInMinutes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.35.0/docs/resources/cloudformation_stack#timeouts CloudformationStack#timeouts}
	Timeouts *CloudformationStackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

