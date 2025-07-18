// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerendpoint


type SagemakerEndpointDeploymentConfigAutoRollbackConfiguration struct {
	// alarms block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.4.0/docs/resources/sagemaker_endpoint#alarms SagemakerEndpoint#alarms}
	Alarms interface{} `field:"optional" json:"alarms" yaml:"alarms"`
}

