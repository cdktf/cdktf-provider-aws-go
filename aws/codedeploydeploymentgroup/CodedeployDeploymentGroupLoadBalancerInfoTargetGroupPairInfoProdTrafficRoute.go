// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.80.0/docs/resources/codedeploy_deployment_group#listener_arns CodedeployDeploymentGroup#listener_arns}.
	ListenerArns *[]*string `field:"required" json:"listenerArns" yaml:"listenerArns"`
}

