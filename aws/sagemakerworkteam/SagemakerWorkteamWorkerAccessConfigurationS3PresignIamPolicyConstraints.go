// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sagemakerworkteam


type SagemakerWorkteamWorkerAccessConfigurationS3PresignIamPolicyConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/sagemaker_workteam#source_ip SagemakerWorkteam#source_ip}.
	SourceIp *string `field:"optional" json:"sourceIp" yaml:"sourceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/5.75.1/docs/resources/sagemaker_workteam#vpc_source_ip SagemakerWorkteam#vpc_source_ip}.
	VpcSourceIp *string `field:"optional" json:"vpcSourceIp" yaml:"vpcSourceIp"`
}

