// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codedeploydeploymentgroup


type CodedeployDeploymentGroupEcsService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/codedeploy_deployment_group#cluster_name CodedeployDeploymentGroup#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.23.0/docs/resources/codedeploy_deployment_group#service_name CodedeployDeploymentGroup#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

