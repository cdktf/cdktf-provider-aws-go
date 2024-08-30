// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupTriggerConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/codedeploy_deployment_group#trigger_events CodedeployDeploymentGroup#trigger_events}.
	TriggerEvents *[]*string `field:"required" json:"triggerEvents" yaml:"triggerEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/codedeploy_deployment_group#trigger_name CodedeployDeploymentGroup#trigger_name}.
	TriggerName *string `field:"required" json:"triggerName" yaml:"triggerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.65.0/docs/resources/codedeploy_deployment_group#trigger_target_arn CodedeployDeploymentGroup#trigger_target_arn}.
	TriggerTargetArn *string `field:"required" json:"triggerTargetArn" yaml:"triggerTargetArn"`
}

