// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupLoadBalancerInfo struct {
	// elb_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/resources/codedeploy_deployment_group#elb_info CodedeployDeploymentGroup#elb_info}
	ElbInfo interface{} `field:"optional" json:"elbInfo" yaml:"elbInfo"`
	// target_group_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/resources/codedeploy_deployment_group#target_group_info CodedeployDeploymentGroup#target_group_info}
	TargetGroupInfo interface{} `field:"optional" json:"targetGroupInfo" yaml:"targetGroupInfo"`
	// target_group_pair_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.31.0/docs/resources/codedeploy_deployment_group#target_group_pair_info CodedeployDeploymentGroup#target_group_pair_info}
	TargetGroupPairInfo *CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfo `field:"optional" json:"targetGroupPairInfo" yaml:"targetGroupPairInfo"`
}

