// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apprunnerservice


type ApprunnerServiceNetworkConfigurationEgressConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/apprunner_service#egress_type ApprunnerService#egress_type}.
	EgressType *string `field:"optional" json:"egressType" yaml:"egressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.49.0/docs/resources/apprunner_service#vpc_connector_arn ApprunnerService#vpc_connector_arn}.
	VpcConnectorArn *string `field:"optional" json:"vpcConnectorArn" yaml:"vpcConnectorArn"`
}

